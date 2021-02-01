package main

import "fmt"

func main() {

    // var declares 1 or more variables
    var a = "initial"
    fmt.Println(a)

    // you can declare multiple variables at onece
    var b, c int = 1, 2
    fmt.Println(b, c)

    // go will infer the type of initialized variables
    var d = true
    fmt.Println(d)

    // variables declared without a corresponding initialization are zero values
    // for example, the zero value for an int is 0
    var e int
    fmt.Println(e)

    // the := syntax is shorthand for declaring and initilizing a variable
    f := "apple"
    fmt.Println(f)


}
