package benchmarks

import (
    "fmt"
    "net"
    "net/http"
    "time"
    "yuzelz-vps-benchmark/utils"
)

type NetworkResult struct {
    DownloadSpeed float64
    Latency       int
}

func RunNetworkBenchmark() NetworkResult {
    fmt.Printf("\n%s🌐 Network Benchmark%s\n", utils.ColorBlue, utils.ColorReset)
    fmt.Println(strings.Repeat("-", 50))

    latency := testLatency()
    downloadSpeed := testDownloadSpeed()

    fmt.Printf("\n%s✅ Network Benchmark Complete%s\n", utils.ColorGreen, utils.ColorReset)

    return NetworkResult{
        DownloadSpeed: downloadSpeed,
        Latency:       latency,
    }
}

func testLatency() int {
    fmt.Print("Testing network latency... ")
    
    start := time.Now()
    conn, err := net.DialTimeout("tcp", "8.8.8.8:80", 5*time.Second)
    if err != nil {
        fmt.Printf("%sFailed%s\n", utils.ColorRed, utils.ColorReset)
        return 0
    }
    defer conn.Close()
    
    elapsed := time.Since(start)
    latency := elapsed.Milliseconds()
    
    fmt.Printf("%s%d ms%s\n", utils.ColorGreen, latency, utils.ColorReset)
    return int(latency)
}

func testDownloadSpeed() float64 {
    fmt.Print("Testing download speed... ")
    
    client := &http.Client{Timeout: 30 * time.Second}
    start := time.Now()
    
    resp, err := client.Get("http://speedtest.tele2.net/10MB.zip")
    if err != nil {
        fmt.Printf("%sFailed%s\n", utils.ColorRed, utils.ColorReset)
        return 0
    }
    defer resp.Body.Close()
    
    data := make([]byte, 1024)
    total := 0
    
    for {
        n, err := resp.Body.Read(data)
        total += n
        if err != nil {
            break
        }
    }
    
    elapsed := time.Since(start).Seconds()
    speed := float64(total) / elapsed / 1024 / 1024 // MB/s
    
    fmt.Printf("%s%.2f MB/s%s\n", utils.ColorGreen, speed, utils.ColorReset)
    return speed
}
