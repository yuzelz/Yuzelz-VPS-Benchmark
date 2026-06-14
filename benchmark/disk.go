package benchmarks

import (
    "fmt"
    "os"
    "time"
    "yuzelz-vps-benchmark/utils"
)

type DiskResult struct {
    WriteSpeed float64
    ReadSpeed  float64
    Iops       int
}

func RunDiskBenchmark() DiskResult {
    fmt.Printf("\n%s💾 Disk Benchmark%s\n", utils.ColorBlue, utils.ColorReset)
    fmt.Println(strings.Repeat("-", 50))

    writeSpeed := testWriteSpeed()
    readSpeed := testReadSpeed()
    iops := testIOPS()

    fmt.Printf("\n%s✅ Disk Benchmark Complete%s\n", utils.ColorGreen, utils.ColorReset)

    return DiskResult{
        WriteSpeed: writeSpeed,
        ReadSpeed:  readSpeed,
        Iops:       iops,
    }
}

func testWriteSpeed() float64 {
    fmt.Print("Testing disk write speed... ")
    
    data := make([]byte, 1024*1024) // 1MB
    filename := "benchmark_test.tmp"
    
    start := time.Now()
    f, _ := os.Create(filename)
    
    total := 0
    for total < 100*1024*1024 { // 100MB
        n, _ := f.Write(data)
        total += n
    }
    f.Sync()
    elapsed := time.Since(start).Seconds()
    f.Close()
    os.Remove(filename)
    
    speed := float64(total) / elapsed / 1024 / 1024 // MB/s
    fmt.Printf("%s%.2f MB/s%s\n", utils.ColorGreen, speed, utils.ColorReset)
    return speed
}

func testReadSpeed() float64 {
    fmt.Print("Testing disk read speed... ")
    
    filename := "benchmark_read.tmp"
    data := make([]byte, 100*1024*1024) // 100MB
    f, _ := os.Create(filename)
    f.Write(data)
    f.Close()
    
    start := time.Now()
    f, _ = os.Open(filename)
    buffer := make([]byte, 1024*1024)
    total := 0
    
    for total < 100*1024*1024 {
        n, _ := f.Read(buffer)
        total += n
    }
    
    elapsed := time.Since(start).Seconds()
    f.Close()
    os.Remove(filename)
    
    speed := float64(total) / elapsed / 1024 / 1024 // MB/s
    fmt.Printf("%s%.2f MB/s%s\n", utils.ColorGreen, speed, utils.ColorReset)
    return speed
}

func testIOPS() int {
    fmt.Print("Testing disk IOPS... ")
    
    filename := "benchmark_iops.tmp"
    f, _ := os.Create(filename)
    
    start := time.Now()
    count := 0
    
    for time.Since(start).Seconds() < 1.0 {
        f.WriteAt([]byte("x"), int64(count*512))
        count++
    }
    
    f.Close()
    os.Remove(filename)
    
    fmt.Printf("%s%d IOPS%s\n", utils.ColorGreen, count, utils.ColorReset)
    return count
}
