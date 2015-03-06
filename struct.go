// Package main provides ...
package main

import (
    "fmt"
)

type person struct {
    id  int
    name string
    age int
}

func main() {
    fmt.Println(person{0, "Alice", 6})
    fmt.Println(person{id: 1, name: "Bob", age: 6})
    fmt.Println(person{name:"Chris"})
    fmt.Println(&person{name:"Chris"})

    p1 := person{3, "Donk", 7}
    fmt.Println("p1 id: ", p1.id)
    
    pp := &p1
    fmt.Println("pp id: ", pp.id)
    pp.age = 6
    fmt.Println("p1 age: ", p1.age)
    fmt.Println("pp age: ", pp.age)

    pp.produce()
    pp.sayId()
    p1.produce()
    p1.sayId()
}

//add method, like define a method in class of python
func (r *person) produce() {
    fmt.Println("Hi, I am", r.name, ", I am", r.age, "years old")
}

func (r person) sayId() {
    fmt.Println("My id is: ", r.id)
}
