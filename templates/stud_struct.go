package main

import (
    "log"
    "text/template"
    "os"
)

type Student struct {
    // exported field since it begins with a capital letter
    Name string
}

func main() {
    //fmt.Println("Hello World")

    // define an instance
    s := Student{"Brendan"}

    // create a new template with some name
    temp := template.New("test")

    // parse some content and generate a template
    temp, err := temp.Parse("Hello {{.Name}}!")
    if err != nil {
        log.Fatal("Parse: ", err)
        return
    }

    // merge template 'temp' with content of 's'
    err1 := temp.Execute(os.Stdout, s)
    if err1 != nil {
        log.Fatal("Execute: ", err1)
        return
    }

}
