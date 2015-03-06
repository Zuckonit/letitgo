// Package main provides ...
package main

import (
    "fmt"
    "errors"
)

type argError struct {
    arg int
    prob string
}

func (e *argError) Error() string{
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f1(arg int) (int, error) {
    if arg == 42 {
        return -1, errors.New("42 is so magic")
    }
    return arg, nil
}

func f2(arg int) (int, error) {
    if arg == 42 {
        return -1, &argError{arg, "is so magic"}
    }
    return arg, nil
}

func main() {
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 failed: ", e)
        } else {
            fmt.Println("f1 success: ", r)
        }
    }

    for _, i := range []int{6, 42} {
        if r, e := f2(i); e !=nil {
            fmt.Println("f2 failed: ", e)
        } else {
            fmt.Println("f2 success: ", r)
        }
    }

    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }

}

