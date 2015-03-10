// Package main provides ...
package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    m := map[string]int{"k1": 1, "k2":2}
    jm, _:= json.Marshal(m)
    //sm := string(jm)

    var dat map[string]interface{}
    json.Unmarshal(jm, &dat)
    fmt.Println(dat)
}
