package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"text/tabwriter"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Docker is not connected: %v", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: false})
	if err != nil {
		log.Fatalf("Can't get Container list: %v", err)
	}

	// Tabwriter setup - Column
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "NAME\tCPU %\tMEM USAGE / LIMIT\tMEM %\tSTATUS")
	fmt.Fprintln(w, "----\t-----\t-----------------\t-----\t------")

	for _, c := range containers {
		stats, err := cli.ContainerStats(ctx, c.ID, false)
		if err != nil {
			continue
		}

		var v types.StatsJSON
		if err := json.NewDecoder(stats.Body).Decode(&v); err != nil {
			stats.Body.Close()
			if err != io.EOF {
				continue
			}
		}
		stats.Body.Close()

		// CPU Calculation
		cpuPercent := calculateCPUPercent(&v)

		// Memory Calculation (MiB)
		memUsageRaw := float64(v.MemoryStats.Usage)
		memLimitRaw := float64(v.MemoryStats.Limit)

		memUsageMiB := memUsageRaw / (1024 * 1024)
		memLimitMiB := memLimitRaw / (1024 * 1024)

		// RAM Percentage
		memPercent := 0.0
		if memLimitRaw > 0 {
			memPercent = (memUsageRaw / memLimitRaw) * 100.0
		}

		name := c.Names[0][1:]

		// Table
		fmt.Fprintf(w, "%s\t%.2f%%\t%.2fMiB / %.2fMiB\t%.2f%%\t%s\n",
			name, cpuPercent, memUsageMiB, memLimitMiB, memPercent, c.Status)
	}
	w.Flush()
}

func calculateCPUPercent(v *types.StatsJSON) float64 {
	var cpuPercent = 0.0
	cpuDelta := float64(v.CPUStats.CPUUsage.TotalUsage) - float64(v.PreCPUStats.CPUUsage.TotalUsage)
	systemDelta := float64(v.CPUStats.SystemUsage) - float64(v.PreCPUStats.SystemUsage)

	if systemDelta > 0.0 && cpuDelta > 0.0 {
		cpuPercent = (cpuDelta / systemDelta) * float64(v.CPUStats.OnlineCPUs) * 100.0
	}
	return cpuPercent
}
