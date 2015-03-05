package main

import (
    "fmt"
    "math"
)

const s = "constant"

func main() {
    fmt.Println(s)

    const hour = 24
    fmt.Println(hour)
    fmt.Println(float64(hour))

    fmt.Println(math.Sin(hour))
}
