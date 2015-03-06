// Package main provides ...
package main

import (
    "fmt"
)

//like pointer in c
func main() {
    i := 1
    fmt.Println("initial value: ",  i)

    zerovalue(i)
    fmt.Println("value after zerovalue: ",  i)

    zeroptr(&i)
    fmt.Println("value after zeroptr: ", i)
    
    fmt.Println("pointer address: ", &i)
}

func zerovalue(val int) {
    val = 0
}

func zeroptr(val *int) {
    *val = 0
}

