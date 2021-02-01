package main

import "fmt"

// the (int, int) in this function signature shows that the 
// function returns 2 ints
func vals() (int, int) {
    return 3, 7
}

func main() {

    // here we use the 2 different return values from the call
    // with multipl assignment
    a, b := vals()
    fmt.Println(a, b)

    // if you only want a subset of the returned values, use the
    // blank identifier _
    _, c := vals()
    d, _ := vals()
    fmt.Println(c, d)
}

