package main

import (
    "fmt"
    "time"
)

func main() {
    // tickers use a similar mechanism to timers: a channel that
    // is sent values. here we'll use the select builtin on the
    // channel to awai the values as they arrive every 500 ms
    ticker := time.NewTicker(500 * time.Millisecond)
    done := make(chan bool)

    go func() {
        for {
            select {
            case <-done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    // tickers can be stopped like timers. once a ticker 
    // is stopped, it won't receive any more values on its channel
    // we'll stop ours after 1600 ms
    time.Sleep(1600 * time.Millisecond)
    ticker.Stop()
    done <- true
    fmt.Println("Ticker stopped")
}
