package main

import (
    "fmt"
    "time"
    "sync"
)

// this is the funciton we'll run in every goroutine. note that
// a waitGroup must be passed to functions by pointer
func worker(id int, wg *sync.WaitGroup) {


    // on return, notify the waitGroup that we're done
    defer wg.Done()

    fmt.Printf("Worker %d starting\n", id)

    // sleep to simulate an expensive task
    time.Sleep(time.Second)
    fmt.Printf("worker %d done\n", id)
}

func main() {
    // this waitGroup is used to wait for all the goroutines
    // launched here to finish
    var wg sync.WaitGroup

    // launch several goroutines and increment the waitGroup
    // counter for each
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }

    // block until the waitGroup counter goes back to 0; all the
    // workers notified they're done
    wg.Wait()
}
