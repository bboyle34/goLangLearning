package main

import "fmt"

type rect struct {
    width, height int
}

// this area method has a receiver type of *rect
func (r *rect) area() int {
    return r.width * r.height
}

// methods can be defined for either pointer or value receiver
// types. here's an example of a value receiver
func (r rect) perim() int {
    return 2*r.width + 2*r.height
}

func main() {
    r := rect{10, 5}

    // here we call the 2 methods defined for our struct
    fmt.Println("area:", r.area())
    fmt.Println("perimeter:", r.perim())

    // go automatically handles conversion between values and
    // pointers for method calls. you may want to use a pointer
    // receiver type to avoid copying on method calls or to allow
    // the method to mutate the receiving struct
    rp := &r
    fmt.Println(rp.area())
    fmt.Println(rp.perim())
}
