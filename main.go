package main

import (
    "fmt"
    "os"
    "yuzelz-vps-benchmark/benchmarks"
    "yuzelz-vps-benchmark/utils"
)

func main() {
    printBanner()
    
    if len(os.Args) > 1 && os.Args[1] == "--quick" {
        runQuickBenchmark()
    } else {
        runFullBenchmark()
    }
    
    printFooter()
}

func printBanner() {
    banner := `
╔══════════════════════════════════════════════════════════╗
║                                                          ║
║     🚀 YUZELZ-VPS-BENCHMARK v1.0                        ║
║     Production-ready VPS Benchmarking Utility           ║
║                                                          ║
╚══════════════════════════════════════════════════════════╝
`
    fmt.Println(utils.Colorize(banner, utils.ColorCyan))
}

func runFullBenchmark() {
    fmt.Println(utils.Colorize("\n🔧 Running FULL Benchmark Suite\n", utils.ColorYellow))
    
    cpuResult := benchmarks.RunCPUBenchmark()
    memResult := benchmarks.RunMemoryBenchmark()
    diskResult := benchmarks.RunDiskBenchmark()
    netResult := benchmarks.RunNetworkBenchmark()
    uptimeResult := benchmarks.RunUptimeCheck()
    
    printResults(cpuResult, memResult, diskResult, netResult, uptimeResult)
}

func runQuickBenchmark() {
    fmt.Println(utils.Colorize("\n⚡ Running QUICK Benchmark Suite\n", utils.ColorYellow))
    
    cpuResult := benchmarks.RunCPUBenchmark()
    diskResult := benchmarks.RunDiskBenchmark()
    
    printQuickResults(cpuResult, diskResult)
}

func printResults(cpu benchmarks.CPUResult, mem benchmarks.MemoryResult, 
                 disk benchmarks.DiskResult, net benchmarks.NetworkResult,
                 uptime benchmarks.UptimeResult) {
    
    fmt.Printf("\n%s📊 BENCHMARK RESULTS%s\n", utils.ColorPurple, utils.ColorReset)
    fmt.Println(strings.Repeat("=", 60))
    
    table := utils.NewTable([]string{"Component", "Metric", "Value", "Rating"})
    
    // CPU Results
    rating := getCPURating(cpu.SingleCore)
    table.AddRow([]string{"CPU", "Single Core", fmt.Sprintf("%.2f", cpu.SingleCore), rating})
    table.AddRow([]string{"CPU", "Multi Core", fmt.Sprintf("%.2f", cpu.MultiCore), rating})
    
    // Memory Results
    memRating := getMemoryRating(mem.WriteSpeed)
    table.AddRow([]string{"Memory", "Write Speed", fmt.Sprintf("%.2f MB/s", mem.WriteSpeed), memRating})
    table.AddRow([]string{"Memory", "Read Speed", fmt.Sprintf("%.2f MB/s", mem.ReadSpeed), memRating})
    
    // Disk Results
    diskRating := getDiskRating(disk.WriteSpeed)
    table.AddRow([]string{"Disk", "Write Speed", fmt.Sprintf("%.2f MB/s", disk.WriteSpeed), diskRating})
    table.AddRow([]string{"Disk", "Read Speed", fmt.Sprintf("%.2f MB/s", disk.ReadSpeed), diskRating})
    table.AddRow([]string{"Disk", "IOPS", fmt.Sprintf("%d", disk.Iops), diskRating})
    
    // Network Results
    netRating := getNetworkRating(net.DownloadSpeed)
    table.AddRow([]string{"Network", "Download", fmt.Sprintf("%.2f MB/s", net.DownloadSpeed), netRating})
    table.AddRow([]string{"Network", "Latency", fmt.Sprintf("%d ms", net.Latency), netRating})
    
    table.Print()
    
    // Overall Score
    overallScore := calculateOverallScore(cpu.SingleCore, mem.WriteSpeed, disk.WriteSpeed, net.DownloadSpeed)
    fmt.Printf("\n%s🏆 OVERALL VPS SCORE: %.1f/10.0%s\n", utils.ColorBold, overallScore, utils.ColorReset)
    printRecommendation(overallScore)
}

func printQuickResults(cpu benchmarks.CPUResult, disk benchmarks.DiskResult) {
    fmt.Printf("\n%s📊 QUICK RESULTS%s\n", utils.ColorPurple, utils.ColorReset)
    fmt.Println(strings.Repeat("=", 50))
    
    fmt.Printf("CPU Score: %s%.2f%s\n", utils.ColorGreen, cpu.SingleCore, utils.ColorReset)
    fmt.Printf("Disk Speed: %s%.2f MB/s%s\n", utils.ColorGreen, disk.WriteSpeed, utils.ColorReset)
    
    overallScore := (cpu.SingleCore + disk.WriteSpeed/100) / 2
    fmt.Printf("\nQuick Score: %.1f/10.0\n", overallScore)
}

func getCPURating(score float64) string {
    switch {
    case score >= 15:
        return utils.Colorize("Excellent", utils.ColorGreen)
    case score >= 10:
        return utils.Colorize("Good", utils.ColorCyan)
    case score >= 5:
        return utils.Colorize("Average", utils.ColorYellow)
    default:
        return utils.Colorize("Poor", utils.ColorRed)
    }
}

func getMemoryRating(speed float64) string {
    switch {
    case speed >= 5000:
        return utils.Colorize("Excellent", utils.ColorGreen)
    case speed >= 3000:
        return utils.Colorize("Good", utils.ColorCyan)
    case speed >= 1000:
        return utils.Colorize("Average", utils.ColorYellow)
    default:
        return utils.Colorize("Poor", utils.ColorRed)
    }
}

func getDiskRating(speed float64) string {
    switch {
    case speed >= 500:
        return utils.Colorize("Excellent", utils.ColorGreen)
    case speed >= 200:
        return utils.Colorize("Good", utils.ColorCyan)
    case speed >= 50:
        return utils.Colorize("Average", utils.ColorYellow)
    default:
        return utils.Colorize("Poor", utils.ColorRed)
    }
}

func getNetworkRating(speed float64) string {
    switch {
    case speed >= 50:
        return utils.Colorize("Excellent", utils.ColorGreen)
    case speed >= 20:
        return utils.Colorize("Good", utils.ColorCyan)
    case speed >= 5:
        return utils.Colorize("Average", utils.ColorYellow)
    default:
        return utils.Colorize("Poor", utils.ColorRed)
    }
}

func calculateOverallScore(cpu, mem, disk, net float64) float64 {
    cpuScore := min(cpu/15*10, 10)
    memScore := min(mem/5000*10, 10)
    diskScore := min(disk/500*10, 10)
    netScore := min(net/50*10, 10)
    
    return (cpuScore*0.3 + memScore*0.2 + diskScore*0.3 + netScore*0.2)
}

func printRecommendation(score float64) {
    fmt.Println("\n💡 Recommendation:")
    switch {
    case score >= 8:
        fmt.Println(utils.Colorize("   Excellent VPS! Suitable for production workloads, hosting, and heavy applications.", utils.ColorGreen))
    case score >= 6:
        fmt.Println(utils.Colorize("   Good VPS. Suitable for most web applications and medium workloads.", utils.ColorCyan))
    case score >= 4:
        fmt.Println(utils.Colorize("   Average VPS. Suitable for light websites and development.", utils.ColorYellow))
    default:
        fmt.Println(utils.Colorize("   Poor performance. Consider upgrading your VPS plan.", utils.ColorRed))
    }
}

func printFooter() {
    fmt.Printf("\n%s✨ Benchmark completed successfully!%s\n", utils.ColorGreen, utils.ColorReset)
}

func min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}
