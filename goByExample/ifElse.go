package main

import "fmt"

func main() {

    for i:= 0; i <= 10; i++ {
        // basic example
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }

        // you can have an if statement without an else
        if i%4 == 0 {
            fmt.Println(i, "is divisible by 4")
        }
    }

    // a statement can precede conditionals; any variables
    // declared in this statement are available in all branches
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
    // you don't need parentheses around conditionals but you do need braces
}
