package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    // here's a non blocking receive. if a value is available on 
    // messages then select will take the <- messages case with 
    // that value. if not it will immediately take the default case
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    default:
        fmt.Println("no message received")
    }

    // a non blocking send works similarly. here msg cannot be
    // sent to the messages channel, because the channel has not
    // buffer and there is no reciver. therefore the default case
    // is selected
    msg := "hi"
    select {
    case messages <- msg:
        fmt.Println("sent message", msg)
    default:
        fmt.Println("no message sent")
    }

    // we can use multiple cases above the default clause to
    // implement a multi way non blcoking select. here we 
    // attempt non blcoking receives on both messages and signals
    select {
    case msg := <-messages:
        fmt.Println("received message", msg)
    case sig := <-signals:
        fmt.Println("received signal", sig)
    default:
        fmt.Println("no activity")
    }
}
