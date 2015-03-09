// Package main provides ...
package main

import (
    "fmt"
    "time"
)

func main() {
    tricker := time.NewTicker(time.Millisecond * 500)

    go func() {
        for t := range tricker.C {
            fmt.Println("Trick at: ", t)
        }
    }()
    
    time.Sleep(time.Millisecond * 2000)
    tricker.Stop()
    fmt.Println("Tricker stopped")
}
