package main

import (
    "log"
    "os"
    "text/template"
)

type Person struct {
    Name string
    Emails []string
}

const temp = `{{$name := .Name}}.
{{range .Emails}}
    Name is {{$name}}, email is {{.}}
{{end}}
`

func main() {
    p := Person{
            "Brendan",
            []string {"bboyle34@gmail.com", "bboyle@mitre.org", "bboyle2006@yahoo.com"},
        }

    t := template.New("Person template")

    t, err := t.Parse(temp)
    if err != nil {
        log.Fatal("Parse: ", err)
        return
    }

    err = t.Execute(os.Stdout, p)
    if err != nil {
        log.Fatal("Execute: ", err)
        return
    }


}
