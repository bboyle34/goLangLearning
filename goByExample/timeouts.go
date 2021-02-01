package main

import (
    "fmt"
    "time"
)

func main() {
    // support we're executing an external call
    // that returns its result on a channel c1 after 2s. note that
    // the channel is buffered, so the send in the goroutine is
    // nonblocking. this is a common pattern to precent
    // goroutine leaks in case the channel is never read
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c1 <- "result is 1"
    }()

    // here's the select implementing a timeout. res := <-c1
    // awaits the result and <-time.After awaits a value to be 
    // sent after the timeout of 1 second. since the select
    // proceeds with the first receive that's ready, we'll take
    // the timeout case if the operation takes more thatn 
    // the allowed 1 second
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("timeout 1")
    }

    // if we allow a longer timeout of 3 seconds, then the recieve
    // from c2 will succeed and we'll print the result
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second):
        fmt.Println("timeout 2")
    }
}
