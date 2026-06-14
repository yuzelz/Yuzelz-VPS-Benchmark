package benchmarks

import (
    "fmt"
    "runtime"
    "time"
    "yuzelz-vps-benchmark/utils"
)

type CPUResult struct {
    SingleCore float64
    MultiCore  float64
    Cores      int
    Model      string
}

func RunCPUBenchmark() CPUResult {
    fmt.Printf("\n%s📊 CPU Benchmark%s\n", utils.ColorBlue, utils.ColorReset)
    fmt.Println(strings.Repeat("-", 50))

    cores := runtime.NumCPU()
    model := getCPUModel()

    fmt.Printf("CPU Model: %s\n", utils.Colorize(model, utils.ColorCyan))
    fmt.Printf("Cores: %d\n", cores)

    singleScore := runSingleCore()
    multiScore := runMultiCore(cores)

    fmt.Printf("\n%s✅ CPU Benchmark Complete%s\n", utils.ColorGreen, utils.ColorReset)

    return CPUResult{
        SingleCore: singleScore,
        MultiCore:  multiScore,
        Cores:      cores,
        Model:      model,
    }
}

func runSingleCore() float64 {
    fmt.Print("Testing single-core performance... ")
    start := time.Now()
    
    result := 0.0
    for i := 0; i < 100000000; i++ {
        result += float64(i) * 1.0000001
    }
    
    elapsed := time.Since(start).Seconds()
    score := 100000000 / elapsed / 10000000
    
    fmt.Printf("%s%.2f%s\n", utils.ColorGreen, score, utils.ColorReset)
    return score
}

func runMultiCore(cores int) float64 {
    fmt.Print("Testing multi-core performance... ")
    start := time.Now()
    
    done := make(chan bool, cores)
    for c := 0; c < cores; c++ {
        go func() {
            result := 0.0
            for i := 0; i < 100000000/cores; i++ {
                result += float64(i) * 1.0000001
            }
            done <- true
        }()
    }
    
    for c := 0; c < cores; c++ {
        <-done
    }
    
    elapsed := time.Since(start).Seconds()
    score := 100000000 / elapsed / 10000000
    
    fmt.Printf("%s%.2f%s\n", utils.ColorGreen, score, utils.ColorReset)
    return score
}

func getCPUModel() string {
    // Simplified - in production would read from /proc/cpuinfo
    return "Virtual CPU (KVM/QEMU)"
}
