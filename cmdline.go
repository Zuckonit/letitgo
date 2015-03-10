// Package main provides ...
package main

import (
    . "fmt" //like from fmt import *
    "os"
    "flag"
)


func main() {
    //argv in python
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    
    Println(argsWithProg)
    Println(argsWithoutProg)

    wordPtr := flag.String("word", "foo", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")
    
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    
    flag.Parse()
    Println("word: ", *wordPtr)
    Println("numb: ", *numbPtr)
    Println("fork: ", *boolPtr)
    Println("svar: ", svar)
    Println("tail: ", flag.Args())
}
