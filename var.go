package main

import (
    "fmt"
)

func main() {
    //give an init value
    var s string = "this is a string"
    t := "this is also a string"
    fmt.Println(s)
    fmt.Println(t)
    
    //unpack value, like python
    var b, c int = 1, 2
    d, e := 1, "string c"

    //h, i := 1 //this is error
    //h := 1, 2 // this is error
    h, _ := 1, 2 //its ok, like python
    fmt.Println(h)

    fmt.Println(b, c)
    fmt.Println(d, e)
    
    //default value
    var bool_v bool
    var int_v  int
    var str_v  string
    var float_v float32
    fmt.Println("bool: ", bool_v, "int: ", int_v, "string:", str_v, "float32", float_v)
}
