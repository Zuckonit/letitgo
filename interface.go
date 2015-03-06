// Package main provides ...
package main

import (
    "fmt"
    "math"
    "reflect"
)

type geometry interface {
    area() float64
    perim() float64
}

type square struct {
    width, heigth float64
}

type cycle struct {
    radius float64
}

func main() {
    s := square{1, 2}
    c := cycle{5}
    measuare(&s)
    measuare(&c)
}

func (s *square) area() float64 {
    return s.width * s.heigth
}

func (s *square) perim() float64 {
    return s.width + s.heigth
}

func (c *cycle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (c *cycle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measuare(g geometry) {
    fmt.Println(reflect.TypeOf(g))
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}
