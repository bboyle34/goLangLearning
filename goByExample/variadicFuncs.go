package main

import "fmt"

// here's a function that will take an arbitrary number of
// ints as arguments
func sum(nums ...int) {
    fmt.Println(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {

    // variadic functions can be called in the usual way with
    // individual arguments
    sum(1, 2, 3, 4)
    sum(234, 2)

    // if you already have multiple args in a slice, apply them to a 
    // variadic function using func(slice...) like this
    nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
    sum(nums...)
}

