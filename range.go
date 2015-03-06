// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    //range, like `enumerate` or `in` in python

    //range a slice
    nums := []int{1, 2, 3}
    for idx, num := range nums {
        fmt.Printf("index: %d, value: %d\n", idx, num)
    }
    
    tds := [][]int{{1,2}, {2}}
    for idx, num := range tds {
        fmt.Printf("index: %d, value: %d\n", idx, num)
    }

    //range a map
    maps := map[string]int{"k1": 1, "k2": 2}
    for k, v := range maps {
        fmt.Printf("key: %s, value: %d\n", k, v)
    }

    //range a string, idx, ascii
    str := "ab"
    for i, c := range str {
        fmt.Println(i, c)
    }
}
