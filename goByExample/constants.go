package main

import (
    "fmt"
    "math"
)

// const declares a constant value
const s string = "constant"

func main() {

    fmt.Println(s)

    // a const statement can appear anywhere a var statement can
    const n = 500000000

    // constant expressions perform arithmetic with arbitrary precision
    const d = 3e20 / n
    fmt.Println(d)

    // a numeric constant has no type until its given sucj as by an explicit conversion
    fmt.Println(int64(d))

    // a number can be given a type by utins it in context that requires one,
    // such as a bariable assignment or function call
    fmt.Println(math.Sin(n))

}

