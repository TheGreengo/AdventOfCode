package main

import (
    "os"
    "fmt"
    "math"
    "sort"
    "bufio"
)

func main() {
    // For simple text manipulation, we use bufio
    // To use a bufio scanner, we need a file object
    file, err := os.Open("D1P1.txt")
    defer file.Close()

    if err != nil {
        panic("Error reading D1P1.txt")
    }

    // Set up our scanner
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lines []string
    var lhs []int
    var rhs []int

    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    var temp1, temp2 int
    for i := range lines {
        fmt.Sscanf(lines[i], "%d %d", &temp1, &temp2)
        lhs = append(lhs, temp1)
        rhs = append(rhs, temp2)
    }

    sort.Slice(rhs, func(p int, q int) bool { return rhs[p] < rhs[q]})
    sort.Slice(lhs, func(p int, q int) bool { return lhs[p] < lhs[q]})
    
    diff := 0
    for i := range rhs {
        diff += int(math.Abs(float64(rhs[i] - lhs[i])))
    }

    fmt.Println(diff)

}
