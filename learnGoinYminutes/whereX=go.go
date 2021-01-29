// single line comment
/*
multi line comment
*/

/*
a build tag is a line comment starting with // +build
and can be executed by go build -tags="foo bar" command
build tags are placed befor ethe package clause near or at the top of the file
followed by a blank linke or other line comments
*/
// +build prod, dev, test

// a package clause starts every source file
// main is a special name declaring an executable rather than a library
package main

// import declaration declares library packages referenced in this file
import (
    "fmt"           // a package in go standard library
    //"io/ioutil"   // implements some I/O utility function
    // m "math"     // math library with local alias m
    // "net/http"   // yes, a web server
    // "os"         // os functions like working with the file system
    // "strconv"    // string conversions
)

func newTopic(arg string) {
    fmt.Println()
    fmt.Println("##########")
    fmt.Prinltn(arg)
    fmt.Println("##########")
    fmt.Println()
}

// a function definition. main is special. it is the entry point for the executable program
// love it or hate it, go uses brace brackets
func main() {
    // println outputs a line to stdout
    // it comes from the package fmt
    fmt.Println("helo")

    // call another function within this package
    beyondHello()
}

// functions have parameters in parentheses
// if there are no parameters, empty parentheses are still required
func beyondHello() {
    newTopic("Variables")
    var x int   // variable declaration. variables must be declared before use
    x = 3       // variable assignment
    // short declarations use := to infer the type, declare, and assign
    y := 4
    sum, prod := learnMultiple(x, y)        // function returns two values
    fmt.Println("sum", sum, "prod", prod)   // simple output
    learnTypes()
}

/*
functions can have parameters and multiple return values
here x and y are the arguments and sum and prod are whats returned
note that x and sum receive the type int
*/
func learnMultiple(x, y int) (sum, prod int) {
    return x + y, x * y     // return two values
}

// some built in types and literals
func learnTypes() {
    newTopic("Types")
    // short declaration usually gives you what you want
    str := "learn go"   // string type

    s2 := `a raw string literal
can include line breaks`    // same string type

    // non ASCII literal. go soure is UTF-8
    g := 'Σ'    // rune type, an alias for int32, holds a unicode code point
    
    f := 3.14195    //float64, an IEEE-754 64-bit floating point number
    c  3 + 4i       // complex128, represented internally with two float64s

    // var syntax with initializers
    var u uint = 7  // unsigned, but implementation dependent size as with int
    var pi float32 = 22. / 7

    // conversion syntax with a short declaration
    n := byte('/n') // byte is an alias for uint8

    // arrays have size fixed at compile time
    var a4 [4] int          // an array of 4 ints, initialized to all 0
    a5 := [...]int{3, 1, 5, 10, 100}    // an array initialized with a fixed size of five
    // elements, with values 3, 1, 5, 10, 100

    // arrays have value semantics
    a4_cpy := a4            // a4_cpy is a copy of a4, two separate instances
    a4_cpy[0] = 25          // only a4_cpy is changed, a4 stays the same
    fmt.Println(a4_cpy[0] == a4[0]) // false

    // slices have dynamic size. arrays and slices have advantages
    // but use cases for slices are much more common

}


















