package main

import "fmt"

// this function intSeq returns another function, which we
// define anonymously in the body of intSeq. the returned
// functions closes over the variable i to form closure
func intSeq() func() int {
    i := 0
    return func() int {
        i ++
        return i
    }
}

func main() {

    // we call intSeq, assigning the result (a function) to nextInt.
    // this function value captuires its own i calue, which will be
    // updated each time we call nextInt
    nextInt := intSeq()

    // see the effect of the closure by calling nextInt a few times
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // to confirm that the state is unique to that particular
    // function, create and test a new one
    newInts := intSeq()
    fmt.Println(newInts())
}
