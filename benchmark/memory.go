package benchmarks

import (
    "fmt"
    "time"
    "yuzelz-vps-benchmark/utils"
)

type MemoryResult struct {
    ReadSpeed   float64
    WriteSpeed  float64
    TotalMemory uint64
}

func RunMemoryBenchmark() MemoryResult {
    fmt.Printf("\n%s📈 Memory Benchmark%s\n", utils.ColorBlue, utils.ColorReset)
    fmt.Println(strings.Repeat("-", 50))

    size := 100 * 1024 * 1024 // 100MB
    data := make([]byte, size)

    writeSpeed := testWriteSpeed(data)
    readSpeed := testReadSpeed(data)

    fmt.Printf("\n%s✅ Memory Benchmark Complete%s\n", utils.ColorGreen, utils.ColorReset)

    return MemoryResult{
        ReadSpeed:  readSpeed,
        WriteSpeed: writeSpeed,
    }
}

func testWriteSpeed(data []byte) float64 {
    fmt.Print("Testing memory write speed... ")
    start := time.Now()
    
    for i := range data {
        data[i] = byte(i % 256)
    }
    
    elapsed := time.Since(start).Seconds()
    speed := float64(len(data)) / elapsed / 1024 / 1024 // MB/s
    
    fmt.Printf("%s%.2f MB/s%s\n", utils.ColorGreen, speed, utils.ColorReset)
    return speed
}

func testReadSpeed(data []byte) float64 {
    fmt.Print("Testing memory read speed... ")
    start := time.Now()
    
    sum := 0
    for i := range data {
        sum += int(data[i])
    }
    _ = sum // prevent optimization
    
    elapsed := time.Since(start).Seconds()
    speed := float64(len(data)) / elapsed / 1024 / 1024 // MB/s
    
    fmt.Printf("%s%.2f MB/s%s\n", utils.ColorGreen, speed, utils.ColorReset)
    return speed
}
