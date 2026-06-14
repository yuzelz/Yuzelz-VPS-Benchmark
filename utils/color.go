package utils

import "runtime"

const (
    ColorReset  = "\033[0m"
    ColorRed    = "\033[31m"
    ColorGreen  = "\033[32m"
    ColorYellow = "\033[33m"
    ColorBlue   = "\033[34m"
    ColorPurple = "\033[35m"
    ColorCyan   = "\033[36m"
    ColorWhite  = "\033[37m"
    ColorBold   = "\033[1m"
)

func DisableColors() {
    if runtime.GOOS == "windows" {
        // Colors disabled on Windows by default
    }
}

func Colorize(text, color string) string {
    if runtime.GOOS == "windows" {
        return text
    }
    return color + text + ColorReset
}
