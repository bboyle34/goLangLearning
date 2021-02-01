package main

import (
    "html/template"
    "os"
)

type Test struct {
    HTML string
    SafeHTML template.HTML
    Title string
    Path string
    Dog Dog
    Map map[string]string
}

type Dog struct {
    Name string
    Age int
}

func main() {
    //fmt.Println("Hello World")
    t, err := template.ParseFiles("context.gohtml")
    if err != nil {
        panic(err)
    }

    data := Test {
        HTML: "<h1>A header!</h1>",
        SafeHTML: template.HTML("<h1>A safe header</h1>"),
        Title: "Backslash! An in depth look at the \"\\\" character.",
        Path: "/dashboard/settings",
        Dog: Dog{"Gordie", 15},
        Map: map[string]string {
                "key": "value",
                "other key": "other value",
                "fuck": "you",
        },
    }

    err = t.Execute(os.Stdout, data)
    if err != nil {
        panic(err)
    }

}
