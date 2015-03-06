// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    //simple func
    pls := plus(1, 2)
    fmt.Println("plus: ", pls)
    
    //multi return
    v1, v2 := values(1, 2)
    fmt.Println("multi return: ", v1, v2)
    
    //multi args
    s := sum(1, 2, 3, 4)
    fmt.Println("multi args: ", s)
    
    //unpack multi args with slice
    args := []int{1, 2, 3, 4}
    fmt.Println("unpack: ", sum(args...))


    //closure
    nint := intSeq()
    for i:=0 ; i<2; i++ {
        fmt.Println(nint())
    }

    //recursion
    rec := fact(10)
    fmt.Println("recursion: ", rec)
}

func plus(a int, b int) int {
    return a+b
}

func values(a int, b int) (int, int) {
    return a, b
}

//like *args in python
func sum(nums ...int) int {
    fmt.Printf("type of nums: %T\n", nums)

    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

//closure
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

//recursion
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n - 1)
}
