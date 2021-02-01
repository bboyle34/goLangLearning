package main

import "fmt"

func main() {

    // unlike arrays, slices are types only by the elements they
    // contain (not the number of elements)
    // to create an empty slice with non zero length, use the built-in make
    // here we make a slice of strings of length 3 (initially zero valued)
    s := make([]string, 3)
    fmt.Println("emp:", s)

    // we can set and get just like with arrays
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // len returns the length of the slice as expected
    fmt.Println("len:", len(s))

    // in addition to these basic operations, slices support several
    // more that make them richer than arrays. one is the builtin
    // append, which returns a slice containing one or more new 
    // value. note that we need to accept a return value from
    // append as we may get a new slice value
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // slices can also be copy'd. here we create an empty slice c
    // of the same length as s and copy into c from s
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // slices support a slice operator with the syntax
    // slice[low:high]. for example, this gets a slice of the
    // elements s[2], s[3], and s[4]
    l := s[2: 5]
    fmt.Println("sl1:", l)

    // this slices up to but not including 5
    l = s[:5]
    fmt.Println("sl2:", l)

    // this slices up from and including 2
    l = s[2:]
    fmt.Println("sl3:", l)

    // we can declare and initialize a variable for slice in
    // a single line as well
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)

    // slices can be composed into multi dimensional data
    // structures. the length of the inner slices can vary, unlike with multi dimensional arrays
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

}
