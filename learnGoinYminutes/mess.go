package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello Brendan, it is", time.Now())
    fmt.Println()

    attempt := 5
    //result = make(chan int, 1)

    go func(att int) {
        att = att + 10
        fmt.Println(att)
    }(attempt)

    attempt = 5 + 5
    fmt.Println(attempt)

    for i := 0; i < 10; i++ {
        fmt.Println(i)
    }
    fmt.Println()

    // i guess we need time.Sleep to wait for the main goroutine 
    // to finish so the concurrent goroutine can end
    time.Sleep(time.Second)

    fmt.Println(attempt)
}
