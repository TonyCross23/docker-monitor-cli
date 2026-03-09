# 🐳 Docker Resource Monitor CLI

A lightweight, real-time command-line tool built with Go to monitor your Docker containers' resource usage (CPU & RAM) with a clean, formatted table output.

## ✨ Features
- **Real-time Stats**: Tracks CPU and Memory usage percentages accurately.
- **Memory Insights**: Displays current memory usage alongside the container's limit (MiB).
- **Clean UI**: Uses `tabwriter` for perfectly aligned, easy-to-read table columns.
- **Cross-Platform**: Ready-to-use binaries for Windows, macOS, and Linux.
- **Single Binary**: No dependencies required—just download and run.

---

## 🚀 Installation

### 🐧 Linux & 🍎 macOS
1. Go to the [Releases](https://github.com/TonyCross23/docker-monitor-cli/releases) page.
2. Download the binary for your OS and architecture (e.g., `linux-amd64` or `darwin-arm64`).
3. Open your terminal and run:
   ```bash
   # Make it executable
   chmod +x docker-monitor-<os>-<arch>

   # Move it to your local bin to use it globally
   sudo mv docker-monitor-<os>-<arch> /usr/local/bin/docker-monitor