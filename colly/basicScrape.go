package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly/v2"
    "os"
)

// writeFile will print any string of text toa file safely by
// checking for erros and syncing at the end
func writeFile(filename string, data []string) error {
    file, err := os.Create(filename)
    if err != nil {
    fmt.Println("There is already a file with this name")
        return err
    }
    defer file.Close()
<<<<<<< HEAD
    fmt.Printf("\n%T\n", file)
=======

>>>>>>> 1ce437176c86b25cc9dc3d3cabe86578e9af6dbc
    fmt.Fprintln(file, "")
    fmt.Fprintln(file, "####################")
    fmt.Fprintln(file, "Links for apnews.com")
    fmt.Fprintln(file, "####################")
    fmt.Fprintln(file, "")
    for _, value := range data {
        //statement := key + " -> " + value
        fmt.Fprint(file, "")
        fmt.Fprintln(file, value)
        fmt.Fprintln(file, "")
        if err != nil {
            fmt.Println(err)
            return err
        }
    }

    return file.Sync()
}

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    // instantiate default collector
    c := colly.NewCollector(
        // maxdepth is 1, so only the links on the scraped page
        // are visited, and not further links are followed
        colly.MaxDepth(3),
        colly.AllowedDomains("apnews.com"),
    )

    var data []string
    count := 0
    // on every element which has href attribute call callback
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        if count > 25 {
            return
        } else {
            link := e.Attr("href")

            // print link to file
            if (link[0] == 'h') {
                fmt.Printf("Link found: %s -> %s\n", e.Text, link)
                fmt.Println()
                data = append(data, link)

                // visit link found on page
                // only those links are visited which are in allowedDomains
                c.Visit(e.Request.AbsoluteURL(link))
                count++
            }
        }
    })

    // before making a request print "Visiting ..."
    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String())
        fmt.Println()
        time.Sleep(1 * time.Second)
    })

    // start scraping on some sites
    c.Visit("https://apnews.com/")
    writeFile("output.txt", data)
}
