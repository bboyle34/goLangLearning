package main

import (
    "fmt"
    "time"
)

func main() {
    // for our example we'll select across two channels
    c1 := make(chan string)
    c2 := make(chan string)
    c10 := make(chan string)

    // each channel will receive a value after some amount of
    // time, to simulate
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "one"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "two"
    }()
    go func() {
        time.Sleep(10 * time.Second)
        c10 <- "ten"
    }()

    // we'll use select to awaut both of these values
    // simultaneously, printing each one as it arrives
    for i := 0; i < 3; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        case msg10 := <-c10:
            fmt.Println("received finally", msg10)
        }
    }
}
