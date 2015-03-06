// Package main provides ...
package main

import (
    "fmt"
)

func main() {
    //define a map
    m := make(map[string]int)
    m["k1"] = 1
    m["k2"] = 2 
    fmt.Println("define then assign map: ", m)

    //define and assign a map
    m2 := map[string]int{"k1": 1, "k2": 2}
    fmt.Println("define and assign map: ", m2)

    //judge whether a key is in a map
    _, prs := m["k2"]
    fmt.Println("k2 in map: ", prs)
    _, prs2 := m["k3"]
    fmt.Println("k3 in map: ", prs2)
    
    //delete a key
    delete(m, "k3")
    fmt.Println("after delete: ", m)
    //remove key
}
