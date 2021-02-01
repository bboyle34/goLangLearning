ackage main

import (
    "html/template"
    "os"
)

func main() {
    //fmt.Println("Hello World")

    t, err := template.ParseFiles("hello.gohtml")
    if err != nil {
        panic(err)
    }

    data := struct {
        Name string
    }{"Brendan Boyle"}

    err = t.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }

}
