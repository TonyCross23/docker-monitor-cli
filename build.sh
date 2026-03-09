# 🐳 Docker Resource Monitor (Go)

A lightweight, real-time command-line tool built with Go to monitor your Docker containers' resource usage (CPU & RAM) with a clean, formatted table output.



## ✨ Features
- **Real-time Monitoring**: Tracks CPU and Memory usage percentages.
- **Memory Insights**: Shows current memory usage vs. container limit (MiB).
- **Clean UI**: Uses `tabwriter` for perfectly aligned table columns.
- **Cross-Platform**: Ready-to-use binaries for Windows, macOS, and Linux.
- **Lightweight**: Single binary with zero dependencies required on the host (except Docker).

---

## 🚀 Quick Start (Installation)

### 🐧 Linux & 🍎 macOS
1. Go to the [Releases](https://github.com/TonyCross23/docker-monitor-cli/releases) page.
2. Download the binary for your architecture (e.g., `linux-amd64` or `darwin-arm64`).
3. Open your terminal and run:
   ```bash
   # Make it executable
   chmod +x docker-monitor-<os>-<arch>

   # Move it to your path to use it globally
   sudo mv docker-monitor-<os>-<arch> /usr/local/bin/docker-monitor