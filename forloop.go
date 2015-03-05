// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    sum, i := 0, 1
    for i <= 100 {
        sum = sum + i
        i = i+1
    }
    fmt.Println(sum)

    s2 := 0
    for j := 0; j <= 100; j++ {
        s2 = s2 + j
    }
    fmt.Println(s2)

    for {
        fmt.Println("forloop")
        break
    }
}
