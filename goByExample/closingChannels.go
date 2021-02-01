package main

import "fmt"

// in this example we'll use a jobs channel to communicate 
// work to be done from the main() goroutine to a worker
// goroutine. when we ahve no more jobs for the worker we'll 
//close the jobs channel
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // here's the worker goroutine. it repeatedly receives from
    // jobs with j, more := <-jobs. in this special 2 value form
    // of receive, the more value will be false if jobs has been 
    // closed and all values in the channel have already been
    // received. we use this to notify on done when we've worked
    // all our jobs
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // this send 3 jobs to the worker over the jobs channel,
    // then closes it
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // we await the worker using the synchronization approach
    // we saw ealier
    <-done
}
