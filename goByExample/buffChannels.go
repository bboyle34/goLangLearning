package main

import "fmt"

func main() {
    // here we make a channel of strings buffering up to 2 values
    messages := make(chan string, 2)

    // because this channel is buffered, we can send these values 
    // into the channel without a corresponding concurrent receive
    messages <- "buffered"
    messages <- "channel"

    // later we can receive these two values as usual
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
