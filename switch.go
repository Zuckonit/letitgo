// Package main provides ...
package main

import (
    "fmt"
    "time"
)

func main() {
    week := time.Now().Weekday()
    switch week {
    case time.Saturday, time.Sunday:
        fmt.Println("it's weekend")
    default:
        fmt.Println("it's a weekday")
    }
}
