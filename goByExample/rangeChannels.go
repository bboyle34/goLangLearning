package main

import "fmt"

// we'll iterate over 2 values in the queue channel
func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // this range iterates over each element as it's received from
    // queue. because we closed the channel above, the iteration
    // terminates after receivinf the 2 elements
    for elem := range queue {
        fmt.Println(elem)
    }

}
