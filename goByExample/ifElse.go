package main

import "fmt"

func main() {
    
    for i:= 0; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Println(i, "is even")
        } else {
            fmt.Println(i, "is odd")
        }
        
        if i%4 == 0 {
            fmt.Println(i, "is divisible by 4")
        }
    }

    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}
