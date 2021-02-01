package main

import "fmt"

func main() {

    // here we use range to sum the numbers in a slice
    // arrays work like this too
    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)

    // range on arrays and slices provides both the index and
    // value for each entry. above we didn't need the index, so
    // we ignored it with the blank identifier _. sometimes we
    // actually want the indexes though
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }

    // range on map iterates over key/value pairs
    kvs := map[string]string{"a":"apple", "b":"banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }

    // range can also iterate over just the keys of a map
    for k := range kvs {
        fmt.Println("key:", k)
    }

    // range on straings iterates over unicode code points.
    // the first value is teh starting byte index of the rune and the
    // second the rune itself
    for i, c := range "go" {
        fmt.Println(i, c)
    }

}
