// Package main provides ...
package main

import (
    "fmt"
)

func f(from string) {
    for i:=0; i<3; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {
    f("direct")
    //like greenlet in python module gevent
    go f("goroutine")
    go func(msg string) {
        fmt.Println(msg)
    }("going")

    var input string
    fmt.Scanln(&input)
}
