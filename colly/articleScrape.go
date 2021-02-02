package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly/v2"
)

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    c := colly.NewCollector(
        colly.AllowedDomains("apnews.com"),
    )

    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        fmt.Printf("Link found: %s -> %s\n", e.Text, link)
        fmt.Println()
        c.Visit(e.Request.AbsoluteURL(link))
    })

    c.OnRequest(func(r *colly.Request) {
        link := r.URL.String()
        if len(link) > 27 {
            if link[0:26] == "https://apnews.com/article" {
                fmt.Printf("%q\n", link[0:26])
                fmt.Println("Visiting", link)
                fmt.Println()
            }
        }
    })

    c.Visit("https://apnews.com/")
}
