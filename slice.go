// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    s := make([]string, 3)
    fmt.Println("make a slice: ", s)
    
    //assign
    s[0] = "a"
    s[1] = "b"
    fmt.Println("assign slice s as: ", s)
    fmt.Println("length of slice s: ", len(s))

    //append
    s = append(s, "c")
    fmt.Println("append slice s as: ", s)

    //print type
    fmt.Printf("type of slice: %T\n", s)

    //slice, like python
    s1 := s[:]
    fmt.Println("slice s1: ", s1)

    s2 := s[1:2]
    fmt.Println("slice s2: ", s2)

    //copy
    s3 := make([]string, len(s))
    copy(s3, s)
    fmt.Println("slice copy of s3: ", s3)

    //address
    printPointer("s", s)
    printPointer("s1", s1)
    printPointer("s3", s3)

    //diff length
    td := make([][]int, 3)
    for i:=0; i < 3; i++ {
        innerLen := i + 1
        td[i] = make([]int, innerLen)
        for j:=0; j<innerLen; j++ {
            td[i][j] = i+j
        }
    }
    fmt.Println("slice td is: ", td)
    
    //define and assign
    dass := [][]int{{1,2}, {1}}
    fmt.Println("define and assign slice dass: ", dass)
}

func printPointer(name string, any interface{}) {
    fmt.Printf("the pointer of %s is: %p\n", name, &any)
}
