package main

import (
    "fmt"
    "time"
)

func main() {
    // timers represent a single event in the future. you tell the
    // timer how long you want to wait, and it provides a channel
    // that will be notified at that time. this timer will wait 2
    // seconds
    timer1 := time.NewTimer(2 * time.Second)

    // the <-time1.C blocks on the timer's channel c until it 
    // sends a value indicating that the timer fired
    <-timer1.C
    fmt.Println("Timer 1 fired")

    // if you just wanted to wait, you have used 
    // time.Sleep. one reason a timer may be useful is that you
    // can cancel the timer before it fires. here's an example of
    // that
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Time 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("timer 2 stopped")
    }

    // give the timer2 enough time to fire, if it ever was going ot,
    // to show it is in fact stopped
    time.Sleep(2 * time.Second)
}
