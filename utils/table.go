package utils

import (
    "fmt"
    "strings"
)

type Table struct {
    headers []string
    rows    [][]string
}

func NewTable(headers []string) *Table {
    return &Table{
        headers: headers,
        rows:    make([][]string, 0),
    }
}

func (t *Table) AddRow(row []string) {
    t.rows = append(t.rows, row)
}

func (t *Table) Print() {
    colWidths := make([]int, len(t.headers))
    for i, header := range t.headers {
        colWidths[i] = len(header)
    }

    for _, row := range t.rows {
        for i, cell := range row {
            if len(cell) > colWidths[i] {
                colWidths[i] = len(cell)
            }
        }
    }

    totalWidth := 0
    for _, w := range colWidths {
        totalWidth += w + 3
    }
    totalWidth += 1

    printLine := func() {
        fmt.Print("+")
        for _, w := range colWidths {
            fmt.Print(strings.Repeat("-", w+2) + "+")
        }
        fmt.Println()
    }

    printLine()
    fmt.Print("|")
    for i, header := range t.headers {
        fmt.Printf(" %-*s |", colWidths[i], header)
    }
    fmt.Println()
    printLine()

    for _, row := range t.rows {
        fmt.Print("|")
        for i, cell := range row {
            fmt.Printf(" %-*s |", colWidths[i], cell)
        }
        fmt.Println()
    }
    printLine()
}
