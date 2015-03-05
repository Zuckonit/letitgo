// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    a := 1
    if a < 1 {
        fmt.Println(a, "is smaller than 1")
    } else if a == 1 {
        fmt.Println(a, "is equal to 1")
    } else {
        fmt.Println(a, "is bigger than 1")
    }
}
