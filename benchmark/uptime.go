package benchmarks

import (
    "fmt"
    "syscall"
    "time"
    "yuzelz-vps-benchmark/utils"
)

type UptimeResult struct {
    Uptime    time.Duration
    LoadAvg   [3]float64
}

func RunUptimeCheck() UptimeResult {
    fmt.Printf("\n%s⏰ System Information%s\n", utils.ColorBlue, utils.ColorReset)
    fmt.Println(strings.Repeat("-", 50))

    var info syscall.Sysinfo_t
    syscall.Sysinfo(&info)
    
    uptime := time.Duration(info.Uptime) * time.Second
    loadAvg := [3]float64{
        float64(info.Loads[0]) / 65536.0,
        float64(info.Loads[1]) / 65536.0,
        float64(info.Loads[2]) / 65536.0,
    }
    
    fmt.Printf("System Uptime: %s\n", utils.Colorize(formatUptime(uptime), utils.ColorCyan))
    fmt.Printf("Load Average: %.2f, %.2f, %.2f\n", loadAvg[0], loadAvg[1], loadAvg[2])
    
    return UptimeResult{
        Uptime:  uptime,
        LoadAvg: loadAvg,
    }
}

func formatUptime(d time.Duration) string {
    days := int(d.Hours() / 24)
    hours := int(d.Hours()) % 24
    minutes := int(d.Minutes()) % 60
    
    if days > 0 {
        return fmt.Sprintf("%d days, %d hours, %d minutes", days, hours, minutes)
    }
    return fmt.Sprintf("%d hours, %d minutes", hours, minutes)
}
