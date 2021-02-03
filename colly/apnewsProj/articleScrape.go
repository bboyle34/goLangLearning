package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly/v2"
    "os"
    "io/ioutil"
    "log"
)

// format how the file will lool
func formatPage(file *os.File, link, h1, para string) {
    result := "\n" + link + "\n----------<h1>" + h1 + "</h1>----------\n<p>" + para + "</p>\n"
    fmt.Fprintln(file, result)
}

// create/append to file
func writeFile(filename string, link string, h1 string, para string, created bool) error {
    // if file is not created, create it with the first line
    if !(created) {
        file, err := os.Create(filename)
        if err != nil {
            return err
        }
        fmt.Fprintln(file, "\n\t\t\t\t\t\t##########\n\t\t\t\t\t\tApnews.com\n\t\t\t\t\t\t##########")
        fmt.Fprintln(file, "\tFormat:\n\tlink\n\t<h1>\n\t<p>\n")
        defer file.Close()
    }
    file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Println(err)
    }
    defer file.Close()
    formatPage(file, link, h1, para)
    return file.Sync()
}

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    // USE THIS VARIABLE TO LIMIT YOUR CRAWL
    crawlTime := 250
    // ANYTHING OVER 1000 WILL TAKE A LONG TIME

    c := colly.NewCollector(
        // only allow links with apnews.com as a domain
        colly.AllowedDomains("apnews.com"),
    )

    // variable to check if file is created or not
    created := false

    // collect the h1 tag from each page
    // only collect 1, some hub pages have multiple
    var h1 string
    var p string
    header := 0
    c.OnHTML("h1", func(e *colly.HTMLElement) {
        if header == 0 {
            h1 = e.Text
        }
        header++
    })

    // collect a p tage from each page
    // most articles have multiple paragraphs, we only need the first
    para := 0
    c.OnHTML("p", func(e *colly.HTMLElement) {
        if para == 0 {
            p = e.Text
        }
        para++
    })

    count := 0
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        // limit total link scrapes to 500 to save time
        if len(link) > 8 && count < crawlTime{
            // make sure our link is an article, that's where we want our data
            if link[1:8] == "article" {
                // reset header and para for each page
                writeFile("output.txt", link, h1, p, created)
                created = true
                header = 0
                para = 0
                h1 = ""
                p = ""
                c.Visit(e.Request.AbsoluteURL(link))
            }
        }
        count++
    })

    c.OnRequest(func(r *colly.Request) {
        link := r.URL.String()
        // on request of each page, print out the URL
        fmt.Println("----------")
        fmt.Println("Visiting", link)
        fmt.Println("----------")
        fmt.Println()
    })

    // initially visit the site
    c.Visit("https://apnews.com/")

    // print out the contents of our file
    data, err := ioutil.ReadFile("output.txt")
    if err != nil {
        log.Fatal(err)
    }
    _ = data

    // use this to print out output.txt, or cat output.txt
    //fmt.Println(string(data))
}
