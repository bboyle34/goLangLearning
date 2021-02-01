package main

import "fmt"

// this person struct type has name and age fields
type person struct {
    name string
    age int
}

// newPerson constructs a new person struct with the given name
// then set's their age to automatically 42
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42

    // you can safely return a pointer to local variable as a local
    // variable will survive the scope of the function
    return &p
}

func main() {
    // this syntax creates a new struct
    fmt.Println(person{"bob", 20})
    // you can name the fields when initializing a struct
    fmt.Println(person{name: "Alice", age: 30})
    // ommitted fields will be zero valued
    fmt.Println(person{name: "Fred"})

    // its idiomatic to enapsulate new struct creation in
    // constructor functions
    fmt.Println(&person{name: "Ann", age: 40})
    fmt.Println(newPerson("Jon"))

    s := person{name: "sean", age: 50}

    // access struct fields with a dot
    // you can also use dots with struct pointers
    // the pointers are automatically dereferenced
    sp := &s
    fmt.Println(sp.age)

    // structus are mutable
    sp.age = 51
    fmt.Println(sp.age)
}
