package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly/v2"
)

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    c := colly.NewCollector()

    // find and visit all links
    c.OnHTML("a[href", func(e *colly.HTMLElement) {
            link := e.Request.AbsoluteURL(e.Attr("href"))
            if link != "" {
                fmt.Println(link)
            }
    })

    c.OnRequest(func(r *colly.Request) {
            fmt.Println("Visiting", r.URL)
    })

    c.Visit("http://apnews.com/")
}
