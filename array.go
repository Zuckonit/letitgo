// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    //define an array
    var a [5]int
    fmt.Println("empty array a: ", a)

    //initial an array
    b := [5]int{1,2,3,4,5}
    fmt.Println("array b", b)

    //length
    fmt.Println("length of array b: ", len(b))
    
    //two-dimensional array
    c := [2][3]int{{1,1,1},{2,2,2}}
    fmt.Println("two-dimensional array c: ", c)

    //string array
    d := [2]string{"a", "b"}
    fmt.Println("string array d: ", d)

    //no initial length
    e := [...]int{1, 2, 3, 4, 5, 6}
    fmt.Println("no initial length array e: ", e)
}
