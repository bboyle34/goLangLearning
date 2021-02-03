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
        //colly.MaxDepth(0),
    )

    headers := 0
    c.OnHTML("h1", func(t *colly.HTMLElement) {
        if headers == 0 {
            fmt.Println("----------", t.Text, "---------")
        }
        headers++
    })

    paragraphs := 0
    c.OnHTML("p", func(p *colly.HTMLElement) {
        if paragraphs == 0 {
            fmt.Println(p.Text)
        }
        paragraphs++
    })

    // depth := 0
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")

        /*
        c.OnHTML("h1", func(t *colly.HTMLElement) {
                fmt.Println(t.Text)
        })
        */
        /*
        c.OnHTML("p", func(p *colly.HTMLElement) {
            if count == 0 {
                fmt.Println(p.Text)
            }
            count++
        })
        */
        //fmt.Println("link found: ", link)
        if len(link) > 27 {
            if link[0:26] == "https://apnews.com/article" {
                c.Visit(e.Request.AbsoluteURL(link))
            }
        }
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println()
        fmt.Println("Visiting:", r.URL.String())
        fmt.Println()
    })

    c.Visit("https://apnews.com/")
    //c.Visit("https://hackerspaces.org/")
}
