package main

import (
    "html/template"
    "net/http"
)

var testTemplate *template.Template

type ViewData struct {
    Name string
}

func main() {
    //fmt.Println("Hello World")

    var err error
    testTemplate, err = template.ParseFiles("hello.gohtml")
    if err != nil {
        panic(err)
    }

    http.HandleFunc("/", handler)
    http.ListenAndServe(":3000", nil)
}
