package main

import (
    "fmt"
    "time"
)

func main() {
    go fmt.Println("Hello from another goroutine")
    fmt.Println("Hello from main goroutine")

    // at this point the program execution stops and all
    // active goroutines are killed


    //let's wait for the other goroutine to finish
    time.Sleep(time.Second)
}
