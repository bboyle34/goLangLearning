package main

import (
    "fmt"
    "time"

)

// for goroutine 1
func Aname() {
    arr1 := [4]string{"rohit", "suman", "aman", "ria"}
    for t1 := 0; t1 <= 3; t1++ {
        time.Sleep(150 * time.Millisecond)
        fmt.Printf("%s\n", arr1[t1])
    }
}

// for goroutine 2
func Aid() {
    arr2 := [4]int{300, 301, 302, 303}
    for t2 := 0; t2 <= 3; t2++ {
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("%d\n", arr2[t2])
    }
}

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    fmt.Println("!...Main Go-routine Start...!")

    // calling goroutine1
    go Aname()

    // calling goroutine2
    go Aid()

    time.Sleep(2000 * time.Millisecond)
    fmt.Println("\n!... Main Go-routine End...!")

}
