package main

import "fmt"

// here's a fimctopm that takes two ints and returns their sum as an int
func plus(a, b int) int {

    // go requires explixit returns, it won't automatically 
    // return the value of the last expression
    return a + b
}

// when you have multiple consecutive paramters of the same type
// you may omit the type name for the like typed parameters
// up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
    return a + b +c
}


func main() {

    // call a function just as you'd expect with name(args)
    res := plus(1, 2)
    fmt.Println("1+2=", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3=", res)

}
