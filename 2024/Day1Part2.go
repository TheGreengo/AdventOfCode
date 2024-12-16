package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    file, err := os.Open("D1P1.txt")

    if err != nil {
        panic("Error opening D1P1.txt")
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    var lhs []int
    var rhs []int
    var temp1, temp2 int

    for scanner.Scan() {
        fmt.Sscanf(scanner.Text(), "%d %d", &temp1, &temp2)
        lhs = append(lhs, temp1)
        rhs = append(rhs, temp2)
    }

    counts := make(map[int]int)

    for i := range rhs {
        if counts[rhs[i]] == 0 {
            counts[rhs[i]] = 1
        } else {
            counts[rhs[i]] += 1
        }
    }

    sum := 0
    for i := range lhs {
        sum += lhs[i] * counts[lhs[i]]
        fmt.Println(lhs[i], " ", counts[lhs[i]], " ", lhs[i] * counts[lhs[i]])
    }

    fmt.Println(sum)
}
