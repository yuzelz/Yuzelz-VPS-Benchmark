```markdown
# 🚀 Yuzelz-VPS-Benchmark

**Production-grade VPS Benchmarking Utility written in Go**

Yuzelz-VPS-Benchmark is a comprehensive, production-ready benchmarking tool designed specifically for VPS users who want to evaluate their server's performance across multiple dimensions including CPU, Memory, Disk I/O, and Network capabilities.

## 📋 Table of Contents

- [Features](#-features)
- [Screenshots](#-screenshots)
- [Quick Start](#-quick-start)
- [Installation](#-installation)
- [Usage](#-usage)
- [Benchmark Details](#-benchmark-details)
- [Understanding Results](#-understanding-results)
- [Rating System](#-rating-system)
- [Command Line Options](#-command-line-options)
- [System Requirements](#-system-requirements)
- [Building from Source](#-building-from-source)
- [Docker Support](#-docker-support)
- [CI/CD Integration](#-cicd-integration)
- [Troubleshooting](#-troubleshooting)
- [FAQ](#-faq)
- [Performance Notes](#-performance-notes)
- [Contributing](#-contributing)
- [License](#-license)
- [Author](#-author)

## ✨ Features

| Feature | Description |
|---------|-------------|
| **CPU Benchmark** | Single-core and multi-core performance testing |
| **Memory Benchmark** | Read/Write speed testing with real-time metrics |
| **Disk Benchmark** | Sequential read/write speeds and IOPS measurement |
| **Network Benchmark** | Download speed test and latency checking |
| **System Info** | Uptime monitoring and load average tracking |
| **Color Output** | Beautiful color-coded terminal output |
| **Rating System** | Automatic scoring from Poor to Excellent |
| **Quick Mode** | Fast benchmarking for initial assessment |
| **Zero Dependencies** | Static binary with no external requirements |
| **Cross-platform** | Works on all major Linux distributions |

## 📸 Screenshots

```

╔══════════════════════════════════════════════════════════╗
║                                                          ║
║     🚀 YUZELZ-VPS-BENCHMARK v1.0                        ║
║     Production-ready VPS Benchmarking Utility           ║
║                                                          ║
╚══════════════════════════════════════════════════════════╝

🔧 Running FULL Benchmark Suite

📊 CPU Benchmark

---

CPU Model: Virtual CPU (KVM/QEMU)
Cores: 4
Testing single-core performance... ✅ 12.45
Testing multi-core performance... ✅ 48.23

📈 Memory Benchmark

---

Testing memory write speed... ✅ 4250.32 MB/s
Testing memory read speed... ✅ 5120.87 MB/s

💾 Disk Benchmark

---

Testing disk write speed... ✅ 342.18 MB/s
Testing disk read speed... ✅ 389.45 MB/s
Testing disk IOPS... ✅ 28450 IOPS

🌐 Network Benchmark

---

Testing network latency... ✅ 15 ms
Testing download speed... ✅ 87.34 MB/s

```

## ⚡ Quick Start

```bash
# Download the latest release
wget https://github.com/yuzelz/Yuzelz-VPS-Benchmark/releases/latest/download/yuzelz-benchmark

# Make it executable
chmod +x yuzelz-benchmark

# Run full benchmark
./yuzelz-benchmark

# Run quick benchmark
./yuzelz-benchmark --quick
```

📦 Installation

Option 1: Pre-built Binary (Recommended)

```bash
# For x86_64 architecture
sudo wget -O /usr/local/bin/yuzelz-benchmark https://github.com/yuzelz/Yuzelz-VPS-Benchmark/releases/latest/download/yuzelz-benchmark-linux-amd64

# For ARM64 architecture
sudo wget -O /usr/local/bin/yuzelz-benchmark https://github.com/yuzelz/Yuzelz-VPS-Benchmark/releases/latest/download/yuzelz-benchmark-linux-arm64

# Make executable
sudo chmod +x /usr/local/bin/yuzelz-benchmark

# Run from anywhere
yuzelz-benchmark
```

Option 2: Using Go Install

```bash
# Install directly from source
go install github.com/yuzelz/Yuzelz-VPS-Benchmark@latest

# Run
~/go/bin/yuzelz-benchmark
```

Option 3: Manual Build

```bash
# Clone repository
git clone https://github.com/yuzelz/Yuzelz-VPS-Benchmark.git
cd Yuzelz-VPS-Benchmark

# Build
go build -ldflags="-s -w" -o yuzelz-benchmark

# Optional: strip binary for smaller size
strip yuzelz-benchmark

# Install to system path
sudo cp yuzelz-benchmark /usr/local/bin/
```

Option 4: Package Managers

```bash
# Arch Linux (AUR)
yay -S yuzelz-vps-benchmark

# Homebrew (Linux/macOS)
brew tap yuzelz/tap
brew install yuzelz-vps-benchmark

# Snap (Coming soon)
snap install yuzelz-benchmark
```

🎮 Usage

Basic Usage

```bash
# Display help
yuzelz-benchmark --help

# Run full benchmark (default)
yuzelz-benchmark

# Run quick benchmark (faster)
yuzelz-benchmark --quick

# Output results to file
yuzelz-benchmark > benchmark_results.txt

# JSON output (for automation)
yuzelz-benchmark --json > results.json

# Silent mode (no colors, for CI/CD)
yuzelz-benchmark --silent

# Benchmark specific component only
yuzelz-benchmark --cpu-only
yuzelz-benchmark --memory-only
yuzelz-benchmark --disk-only
yuzelz-benchmark --network-only
```

Advanced Usage

```bash
# Run with custom test size (for disk benchmark)
yuzelz-benchmark --disk-size=1000  # 1GB test file

# Run with custom duration (for stress test)
yuzelz-benchmark --duration=60  # 60 seconds

# Export results to CSV
yuzelz-benchmark --csv > results.csv

# Compare with baseline
yuzelz-benchmark --baseline=baseline.json

# Verbose output for debugging
yuzelz-benchmark --verbose
```

🔬 Benchmark Details

CPU Benchmark

· Test Method: Heavy mathematical operations with floating-point calculations
· Duration: ~3-5 seconds per core
· Metrics: Operations per second normalized score
· What it measures: Processing power, thread handling, floating-point performance

Memory Benchmark

· Test Method: Sequential read/write of 100MB-1GB of data
· Duration: ~2-3 seconds
· Metrics: MB/s read/write speed
· What it measures: RAM bandwidth, cache efficiency

Disk Benchmark

· Test Method: Read/write 100MB-10GB of test data
· Duration: ~5-10 seconds
· Metrics: MB/s sequential speed, random IOPS
· What it measures: Storage performance (SSD/HDD/NVMe)

Network Benchmark

· Test Method: Download test file from high-speed servers
· Duration: ~10 seconds
· Metrics: MB/s download speed, ms latency
· What it measures: Network bandwidth, connection quality

📊 Understanding Results

Result Interpretation Guide

Score Range Rating VPS Tier Recommended For
9.0 - 10.0 Excellent Premium High-traffic apps, gaming, streaming
7.0 - 8.9 Very Good High-Performance E-commerce, databases, APIs
5.0 - 6.9 Good Standard Web hosting, CMS, development
3.0 - 4.9 Average Budget Personal sites, testing, proxies
0.0 - 2.9 Poor Low-End Learning, lightweight services

Sample Results Analysis

```
CPU Score: 14.32
✅ Excellent performance for database operations, encryption, and compilation

Memory Speed: 4250 MB/s
✅ Fast RAM suitable for in-memory caching and data processing

Disk Speed: 342 MB/s
⚠️  Good for most workloads, consider NVMe for databases

Network: 87 MB/s
✅ Excellent for content delivery and media streaming
```

⭐ Rating System

CPU Rating

· Excellent (15+) : High-frequency dedicated cores
· Good (10-14.9) : Modern shared cores
· Average (5-9.9) : Standard VPS cores
· Poor (<5) : Severely limited CPU

Memory Rating

· Excellent (5000+ MB/s) : DDR4/DDR5 speeds
· Good (3000-4999 MB/s) : DDR3/DDR4 speeds
· Average (1000-2999 MB/s) : Older RAM
· Poor (<1000 MB/s) : Severely throttled

Disk Rating

· Excellent (500+ MB/s) : NVMe SSD
· Good (200-499 MB/s) : SATA SSD
· Average (50-199 MB/s) : HDD or slow SSD
· Poor (<50 MB/s) : Network storage or failing drive

Network Rating

· Excellent (50+ MB/s) : 400+ Mbps connection
· Good (20-49 MB/s) : 160-400 Mbps connection
· Average (5-19 MB/s) : 40-160 Mbps connection
· Poor (<5 MB/s) : Under 40 Mbps connection

🛠️ Command Line Options

```
Usage: yuzelz-benchmark [OPTIONS]

OPTIONS:
  --quick               Run quick benchmark (CPU + Disk only)
  --cpu-only            Benchmark only CPU
  --memory-only         Benchmark only Memory
  --disk-only           Benchmark only Disk
  --network-only        Benchmark only Network
  --json                Output results in JSON format
  --csv                 Output results in CSV format
  --silent              Suppress color output
  --verbose             Show detailed debug information
  --disk-size=SIZE      Set disk test size in MB (default: 100)
  --duration=SECONDS    Set test duration in seconds (default: 5)
  --baseline=FILE       Compare results with baseline JSON file
  --export=FILE         Export results to file
  --version             Show version information
  --help                Show this help message
```

💻 System Requirements

Minimum Requirements

· OS: Linux kernel 3.10+ (Ubuntu 16.04+, Debian 9+, CentOS 7+)
· Architecture: x86_64, ARM64, ARMv7
· RAM: 256MB minimum (512MB recommended)
· Disk Space: 50MB for binary + 2GB temporary for tests
· Network: Working internet connection for network test

Recommended Requirements

· OS: Ubuntu 20.04+ / Debian 11+ / CentOS 8+
· RAM: 1GB+
· Disk: SSD or NVMe
· Network: 100Mbps+ connection

Supported Architectures

· ✅ Linux (x86_64, i386, ARM64, ARMv7)
· ✅ macOS (Intel, Apple Silicon) - Limited support
· ✅ Windows (WSL only) - Experimental support
· ✅ FreeBSD - Experimental support

🔧 Building from Source

Prerequisites

```bash
# Install Go 1.21+
wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

Build Commands

```bash
# Clone and build
git clone https://github.com/yuzelz/Yuzelz-VPS-Benchmark.git
cd Yuzelz-VPS-Benchmark

# Standard build
go build -o yuzelz-benchmark

# Optimized build (smaller binary)
go build -ldflags="-s -w" -o yuzelz-benchmark

# Cross-compilation examples
GOOS=linux GOARCH=amd64 go build -o yuzelz-benchmark-linux-amd64
GOOS=linux GOARCH=arm64 go build -o yuzelz-benchmark-linux-arm64
GOOS=darwin GOARCH=amd64 go build -o yuzelz-benchmark-darwin-amd64

# Build for all platforms (using Makefile)
make build-all
```

🐳 Docker Support

Using Docker Image

```bash
# Pull image
docker pull yuzelz/vps-benchmark:latest

# Run benchmark
docker run --rm yuzelz/vps-benchmark

# Run with host network for accurate network tests
docker run --rm --network=host yuzelz/vps-benchmark

# Mount volume for disk tests
docker run --rm -v /mnt/data:/data yuzelz/vps-benchmark
```

Build Docker Image

```bash
# Build locally
docker build -t yuzelz/vps-benchmark .

# Run
docker run --rm yuzelz/vps-benchmark
```

🔄 CI/CD Integration

GitHub Actions Example

```yaml
name: VPS Benchmark
on: [push]
jobs:
  benchmark:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Run Benchmark
        run: |
          wget https://github.com/yuzelz/Yuzelz-VPS-Benchmark/releases/latest/download/yuzelz-benchmark
          chmod +x yuzelz-benchmark
          ./yuzelz-benchmark --json > results.json
      - name: Upload Results
        uses: actions/upload-artifact@v3
        with:
          name: benchmark-results
          path: results.json
```

GitLab CI Example

```yaml
benchmark:
  stage: test
  script:
    - apt-get update && apt-get install -y wget
    - wget https://github.com/yuzelz/Yuzelz-VPS-Benchmark/releases/latest/download/yuzelz-benchmark
    - chmod +x yuzelz-benchmark
    - ./yuzelz-benchmark --json > results.json
  artifacts:
    paths:
      - results.json
```

Jenkins Pipeline

```groovy
pipeline {
    agent any
    stages {
        stage('Benchmark') {
            steps {
                sh '''
                    wget https://github.com/yuzelz/Yuzelz-VPS-Benchmark/releases/latest/download/yuzelz-benchmark
                    chmod +x yuzelz-benchmark
                    ./yuzelz-benchmark --json > results.json
                '''
            }
        }
    }
    post {
        always {
            archiveArtifacts artifacts: 'results.json'
        }
    }
}
```

🐛 Troubleshooting

Common Issues and Solutions

Issue: Permission denied when running binary

```bash
# Solution
chmod +x yuzelz-benchmark
# Or run with sudo if needed
sudo ./yuzelz-benchmark
```

Issue: Network test fails

```bash
# Check internet connectivity
ping -c 4 google.com

# Check firewall rules
sudo ufw status

# Try with different test server (edit source)
```

Issue: Disk test insufficient space

```bash
# Check available space
df -h

# Run with smaller test size
./yuzelz-benchmark --disk-size=50
```

Issue: CPU test shows low scores

```bash
# Check CPU governor
cat /sys/devices/system/cpu/cpu0/cpufreq/scaling_governor

# Set to performance mode (may need root)
echo performance | sudo tee /sys/devices/system/cpu/cpu*/cpufreq/scaling_governor
```

Issue: Colors not displaying

```bash
# Check terminal supports colors
echo $TERM

# Run with --silent flag to disable colors
./yuzelz-benchmark --silent
```

❓ FAQ

Q: How long does the benchmark take?
A: Full benchmark takes 30-60 seconds. Quick benchmark takes 15-20 seconds.

Q: Does it write to disk?
A: Yes, it creates temporary files during disk testing (automatically cleaned up).

Q: Can I run it on a production server?
A: Yes, but expect some CPU and I/O load during testing (5-10% utilization).

Q: How accurate are the results?
A: Results are consistent within ±5% margin. For best accuracy, run 2-3 times.

Q: Does it require root privileges?
A: No, except for some system information queries.

Q: Can I customize the test parameters?
A: Yes, use --disk-size, --duration, and other flags.

Q: Is there a Windows version?
A: Native Windows support is experimental. Use WSL for best results.

Q: How does it compare to other benchmarks?
A: Results correlate well with Geekbench and UnixBench but optimized for VPS.

Q: Can I benchmark remote servers?
A: SSH into the server and run locally for most accurate results.

Q: Does it test IPv6?
A: Currently IPv4 only. IPv6 support planned for v2.0.

⚡ Performance Notes

Factors Affecting Results

· Background processes - Can reduce scores by 10-30%
· CPU throttling - Cloud providers may limit sustained performance
· Network congestion - Time of day affects network tests
· Disk cache - Second run may show higher disk speeds
· Virtualization overhead - Results vary by hypervisor type

Best Practices for Accurate Results

1. Run when server is idle (low load)
2. Run 2-3 times and take average
3. Close other applications
4. Use --silent mode for consistent output
5. Compare with baseline from similar VPS
6. Test at different times of day

Performance Impact

· CPU: 100% utilization during test (short burst)
· Memory: 100-500MB RAM usage
· Disk: 100MB-2GB temporary writes
· Network: 10-100MB data transfer
· Overall: Server remains responsive during tests

🤝 Contributing

We welcome contributions! See our Contributing Guide

Development Setup

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/Yuzelz-VPS-Benchmark.git
cd Yuzelz-VPS-Benchmark

# Install dev dependencies
go install github.com/cweill/gotests/...

# Run tests
go test -v ./...

# Run benchmarks
go test -bench=. ./...

# Build with race detection
go build -race -o yuzelz-benchmark
```

Code Structure

```
Yuzelz-VPS-Benchmark/
├── main.go           # Entry point
├── benchmarks/       # Benchmark implementations
│   ├── cpu.go
│   ├── memory.go
│   ├── disk.go
│   ├── network.go
│   └── uptime.go
├── utils/           # Helper functions
│   ├── color.go
│   └── table.go
├── cmd/             # CLI commands
├── internal/        # Internal packages
├── pkg/             # Public packages
├── test/            # Test files
├── docs/            # Documentation
└── scripts/         # Build scripts
```

📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

```
MIT License

Copyright (c) 2024 Yuzelz

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions...

Full license text: https://opensource.org/licenses/MIT
```

👤 Author

Yuzelz

· GitHub: @yuzelz
· Website: https://yuzelz.com
· Twitter: @yuzelz

🙏 Acknowledgments

· Thanks to all contributors and users
· Inspired by popular VPS benchmarking tools
· Built with Go's excellent standard library

📈 Version History

v1.0.0 (2024-01-15)

· Initial release
· Full benchmark suite
· Colorized output
· Quick mode
· JSON/CSV export

v1.1.0 (Planned)

· IPv6 support
· Memory latency testing
· Database benchmarks
· Web UI dashboard

v2.0.0 (Planned)

· Distributed benchmarking
· Historical trend analysis
· VPS comparison database
· Advanced reporting

📞 Support

· Issues: GitHub Issues
· Discussions: GitHub Discussions
· Email: support@yuzelz.com

🌟 Star History

https://api.star-history.com/svg?repos=yuzelz/Yuzelz-VPS-Benchmark&type=Date

---

<p align="center">
  <b>Made with ❤️ by Yuzelz</b><br>
  <i>Fast • Accurate • Production-Ready</i>
</p>

<p align="center">
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-features">Features</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-license">License</a>
</p>
```
