// Package main provides ...
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second * 1)
    fmt.Println("done")
    done <- true
}


//name chan <- type  send
//name <- chan type  receive
func ping(pings chan<- string, msg string) {
    pings <- msg
}

func pong(pings<-chan string, pongs chan <- string) { 
    msg := <-pings
    pongs <- msg
}

//select, like joinall in gevent
func test_select() {
    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        time.Sleep(time.Second*1)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second*2)
        c2 <- "two"
    }()

    for i:=0; i<2; i++ {
        select {
        case m1:= <-c1:
            fmt.Println("received", m1)
        case m2:= <-c2:
            fmt.Println("received", m2)
        }
    }
}

func main() {
    //make(chan val-type)
    msg := make(chan string, 2)
    f := func(m string) {msg <- m}
    go f("ping")
    go f("pong")
    go f("duang")
    fmt.Println(<-msg)
    fmt.Println(<-msg)
    fmt.Println(<-msg)
    
    //sync
    done := make(chan bool, 1)
    go worker(done)

    <-done

    //ping pong
    pings := make(chan string, 1)
    pongs := make(chan string, 1)

    ping(pings, "ping")
    pong(pings, pongs)
    fmt.Println(<-pongs)

    test_select()
}
