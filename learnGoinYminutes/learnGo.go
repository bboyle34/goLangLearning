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
    "io/ioutil"   // implements some I/O utility function
    m "math"     // math library with local alias m
    "net/http"   // yes, a web server
    "os"         // os functions like working with the file system
    "strconv"    // string conversions
    "time"
    "math/rand"
)

func newTopic(arg string) {
    fmt.Println()
    fmt.Println("##########")
    fmt.Println(arg)
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
    c := 3 + 4i       // complex128, represented internally with two float64s

    // var syntax with initializers
    var u uint = 7  // unsigned, but implementation dependent size as with int
    var pi float32 = 22. / 7

    // conversion syntax with a short declaration
    n := byte('\n') // byte is an alias for uint8

    // arrays have size fixed at compile time
    var a4 [4] int          // an array of 4 ints, initialized to all 0
    a5 := [...]int{3, 1, 5, 10, 100}    // an array initialized with a fixed size of five
    // elements, with values 3, 1, 5, 10, 100

    // arrays have value semantics
    newTopic("Arrays")
    a4_cpy := a4            // a4_cpy is a copy of a4, two separate instances
    a4_cpy[0] = 25          // only a4_cpy is changed, a4 stays the same
    fmt.Println(a4_cpy[0] == a4[0]) // false

    // slices have dynamic size. arrays and slices have advantages
    // but use cases for slices are much more common
    newTopic("Slices")
    s3 := []int{4, 5, 9}    // compare to a5. No ellipsis here
    s4 := make([]int, 4)    // allocates slice of 4 ints, initialized to all 0
    var d2 [][]float64      // declaration only, nothing allocated here
    bs := []byte("a slice") // type conversion syntax

    // slices (as well as maps and channels have reference semantics
    s3_cpy := s3            // both variables point to the same instance
    s3_cpy[0] = 0           // which means both are updated
    fmt.Println(s3_cpy[0] == s3[0]) // true

    // because they are dynamic, slices can be appended to on demand
    // to append elements to a slice, the buitl in append() function is used
    // first argument is a slice to which we are appending. commonly,
    // the array variable us updated inplace, as in example below
    s := []int{1, 2, 3}     // result is a slice of length 3
    s = append(s, 4, 5, 6)  // added 3 elements. slice now has length of 6
    fmt.Println(s)          // updated slice is now [1 2 3 4 5 6]
    // to append another slice, instead of list of atomic elements we can
    // pass a reference to a slice or a slice literal like this, with a
    // trailing ellipsis, meaning take a slice and unpack its elements
    // appending them to slice s
    s = append(s,[]int{7, 8, 9}...) // second argument is a slice literal
    fmt.Println(s)      // updated slice is now [1 2 3 4 5 6 7 8 9]

    newTopic("Pointers")
    p, q := learnMemory()   // declares p, 1 to be type pointer to int
    fmt.Println(*p, *q)     // * follows a pointer. this prints two ints

    newTopic("Maps")
    // maps are dynamically growable associative array type, like the 
    // hash or dictionary types of some other languages
    m := map[string]int{"three": 3, "four":4}
    m["one"] = 1

    // unused variables are an error in go
    //the underscore lets you use a variable but discard its vale
    _, _, _, _, _, _, _, _, _, _ = str, s2, g, f, u, pi, n, a5, s4, bs
    // usually you use it to ignore one of the return values of a function
    // for example, in a quick and dirty script you might ifnore the 
    // error value returned from os.Create, and expect that the file
    // will always be created
    file, _ := os.Create("output.txt")
    fmt.Fprint(file, "This is how you create a file, \nby the way")
    file.Close()

    // output of course counts as using a variable
    fmt.Println(s, c, a4, s3, d2, m)

    learnFlowControl()  // back in the flow
}

// go is fully garbage collected. it has pointers but no pointer arithmetic
// you can make a mistake with a nil pointer, but not by incrementing a pointer
// unlike in C/Cpp taking and returning an address of a local variable is also safe
func learnMemory() (p, q *int) {
    // named return values p and q have type pointer to int
    p = new(int) // built in function new allocated memory
    // the allocated int slice is initialized to 0, p is no longer nil
    s := make([]int, 20)    // allocate 20 ints as a single block of memory
    s[3] = 7                // assign one of them
    r := -2                 // declare another local variable
    return &s[3], &r        // & takes the address of an object
}

// it is possible, unlike in many other languages for functions in go
// to have named return values
// assigning a name to the type being returned in the function declaration line
// allows us to easily return from multiple points in a function as well as to
// only use the return keyword, without anything further
func learnNamedReturns(x, y int) (z int) {
    z = x * y
    return  // z is implicit here, because we named it earlier
}

func expensiveComputation() float64 {
    return m.Exp(10)
}

func learnFlowControl() {
    newTopic("Conditionals")
    // if statements require brace brackets, and do not require parenthese
    if true {
        fmt.Println("told ya")
    }
    // formatting is standardized by the command line "go fmt"
    if false {
        // pout
    } else {
        // gloat
    }
    // use switch in preference to chained if statements
    x := 42.0
    switch x {
    case 0:
    case 1, 2: // can have multiple matches on one case
    case 42:
        // cases don't fall through
        /*
        there is a fall through keyword however
        */
    case 43:
        // unreached
    default:
        // default case is optional
    }

    // type switch allows switching on the type of something instead of value
    var data interface{}
    data = ""
    switch c := data.(type) {
    case string:
        fmt.Println(c, "is a string")
    case int64:
        fmt.Printf("%d us ab int64\n", c)
    default:
        // all other cases
    }

    newTopic("For Loop")
    // like if, for doesn't use parens either
    // variables decalred in for and if are local to their scope
    for x := 0; x < 3; x++ {    // ++ is a statement
        fmt.Println("Iteration", x)
    }
    // x == 42 here

    // for is the only loop statement in Go, but it has alternate forms
    for {   // infinite loop
        break       // just kidding
        continue    // unreached
    }

    // you can use range to iterate over an array, a slice, a string, a map, or a channel
    // range returns one (channel) or two values (array, slice, string, and map)
    for key, value := range map[string]int{"one": 1, "two": 2, "three": 3} {
        // for each pair in the map, print key and value
        fmt.Printf("key=%s, value%d\n", key, value)
    }
    // if you only need the value, use the underscore as the key
    for _, name := range []string{"Bob", "Bill", "Joe"} {
        fmt.Printf("Hello, %s\n", name)
    }

    // as with for, := in an if statement means to declare and assign
    // y first, then test y > x
    if y := expensiveComputation(); y > x {
        x = y
    }
    // function literals are closures
    xBig := func() bool {
        return x > 10000    // references x decalred above switch statement
    }
    x = 99999
    fmt.Println("xBig:", xBig())    // true
    x = 1.3e3                       // this makes x == 1300
    fmt.Println("xBig:", xBig())    // false now

    // what's more is function literals may be defined and called inline,
    // acting as an argument to function, as long as:
    // a) function literal is called immediately (),
    // b) result type matches expected type of argument
    fmt.Println("Add + double two numbers: ",
        func(a, b int) int {
            return (a + b) * 2
        }(10, 2))   // called with args 10 and 2
    // => add and double two numbers: 24

    // when you need it, you'll love it
    goto love
love:

    learnFunctionFactory()    // func returning func is fun(3)(3)
    learnDefer()            // a quick detour to an important keyword
    learnInterfaces()       // good stuff coming up
}

func learnFunctionFactory() {
    // next two are equivalent with second being more practive
    newTopic("Functions")
    fmt.Println(sentenceFactory("summer")("a beautiful", "day!"))

    d := sentenceFactory("summer")
    fmt.Println(d("A beautiful", "day!"))
    fmt.Println(d("A lazy", "afternoon"))
}

// decorators are common in other languages. same can be done in go 
// with function literals that accept arguments
func sentenceFactory(mystring string) func(before, after string) string {
    return func(before, after string) string {
        return fmt.Sprintf("%s %s %s", before, mystring, after) // new string
    }
}

func learnDefer() (ok bool) {
    newTopic("Defer")
    // a defer statement pushes a functin call onto a list. the list of saved
    // calls is executed after the surrounding functino returns
    defer fmt.Println("deferred statements execute in reverse order. LIFO")
    defer fmt.Println("\nThis line is being printed first because")
    // defer is commonly used to close a file, so the function closing the 
    // file stays close to the function opening the file
    return true
}

// define stringer as an interface type with one method, string
type Stringer interface {
    String() string
}

// define pair as a struct with two field, ints named x and y
type pair struct {
    x, y int
}

// define a method on type pair. pair now implements stringer because pair has defined
// all the methods in the interface
func (p pair) String() string { // p is called the receiver
    // Sprintf is another public function in package fmt
    // dot syntax references fields of p
    return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

func learnInterfaces() {
    newTopic("Interfaces")
    // brace syntax is a struct literal. it evaluates to an initialized
    // struct. the := syntax declares and initializes p to this struct
    p := pair{3, 4}
    fmt.Println(p.String()) // call string method of p, of type pair
    var i Stringer          // declares i of interface type Stringer
    i = p                   // valid because pair implements stringer
    // call string method of i, of type Stringer. output same as above
    fmt.Println(i.String())

    // functions in the fmt package call the string method to ask an object
    // for a printable representation of itself
    fmt.Println(p)  // output same as above. Println calls string method
    fmt.Println(i)  // output same as above

    learnVariadicParams("great", "learning", "here")
}

// functions can have variadic parameters
func learnVariadicParams(myStrings ...interface{}) {
    newTopic("Parameters")
    // iterate each value of the variadic
    // the underbar here is ignoring the index argument of the array
    for _, param := range myStrings {
        fmt.Println("param:", param)
    }

    // pass variadic value as a variadic parameter
    fmt.Println("params:", fmt.Sprintln(myStrings...))

    learnErrorHandling()
}

func learnErrorHandling() {
    newTopic("Errors")
    // ", ok" idiom used to tell if something worked or notw
    m := map[int]string{3: "three", 2: "two", 1: "one"}
    if x, ok := m[4]; !ok { // ok will be false because 4 is not in the map
        fmt.Println("no one there")
    } else {
        fmt.Print(x)    // x would be the value if it were in the map
    }
    // an error value communicates not just "OK" but more about the problem
    if _, err := strconv.Atoi("non-int"); err != nil {  // _ discard value
        // prints 'strconv.ParseInt: parsing "non-int": invalid syntax'
        fmt.Println(err)
    }
    // we'll revisit interfaces a litter later
    learnConcurrency()
}

// c is a channel, a concurrency safe communicate object
func inc(i int, c chan int) {
    c <- i + 1  // <- is the send operator when a channel appears on the left
}

// we will use inc to increment some numbers concurrently
func learnConcurrency() {
    newTopic("Channels")
    rand.Seed(time.Now().UnixNano())
    // same make function used earlier to make a slice. make allocates and
    // initilizes slices, maps, and channels
    c := make(chan int)
    // start three concurrent goroutines. numbers will be incremented
    // concurrently, perhaps in parallel if the machine is capacble and
    // properly configured. all three send to the same channel
    go inc(0, c)    // go is a statement that starts a new goroutine
    go inc(10, c)
    go inc(-805, c)
    // read three results from the channel and print them out
    // there is no telling in what order the results will arrive
    fmt.Println(<-c, <-c, <-c)  // channel on right, <- is the receive operator

    cs := make(chan string)         // another channel, this one handles strings
    ccs := make(chan chan string)   // a channgel of string channels
    go func() { c <- 84}()          // start a new goroutine just to send a value
    go func() { cs <- "wordy"}()    // again, for cs this time
    // select has syntax like a switch statement but each case involves
    // a channel operation. it selects a case at random out of the cases
    // that are ready to communicate
    select {
    case i := <-c:  // the value received can be assigned to a variable
        fmt.Printf("it's a %T", i)
    case <-cs:      // or the value received can be discarded
        fmt.Println("it's a string")
    case <-ccs:     // empty channel, not ready for communication
        fmt.Println("didn't happen")
    }
    // at this point a value was taken from either c or cs. one of the two
    // goroutines started above has completed, the other will remain blocked

    learnWebProgramming()   // go does it. you want to too
}

// a single function from package http starts a web server
func learnWebProgramming() {

    // first parameter of ListenAndServer is TCP address to listen to
    // second paramtere is an interface, specifically http.Handler
    go func() {
        err := http.ListenAndServe(":8080", pair{})
        fmt.Println(err)    // don't ignore errors
    }()

    requestServer()
}

// make pair an http.Handler by implementing its only method, ServeHTTP
func (p pair) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // serve data with a method of http.ResponseWriter
    w.Write([]byte("You learned Go in Y minutes"))
}

func requestServer() {
    resp, err := http.Get("http://localhost:8080")
    fmt.Println(err)
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("\nWebserver said: `%s`\n", string(body))
}
