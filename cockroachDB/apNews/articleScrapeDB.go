package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly/v2"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

type Article struct {
    articleid int
    link string
    header string
    paragraph string
}

// insert into database
func insertDB(articleid int, link, header, paragraph string) {
    db, err := sql.Open("postgres", "postgresql://root@localhost:26257/apnews_db?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database/insert: ", err)
    }

    if _, err := db.Exec(
        `insert into article (articleid, link, header, paragraph) 
        values ($1, $2, $3, $4);`, articleid, link, header, paragraph); err != nil {
        log.Fatal("Error in sql insert statement: ", err)
    }
}

//get max articleid from database
func maxID() int {
    db, err := sql.Open("postgres", "postgresql://root@localhost:26257/apnews_db?sslmode=disable")
    if err != nil {
        log.Fatal("Error connecting to the database/select: ", err)
    }

    rows, err := db.Query("select max(articleid) from article;")
    if err != nil {
        log.Fatal("Error in sql select statement: ", err)
    }
    defer rows.Close()
    var maxArticleid int
    for rows.Next() {
        if err := rows.Scan(&maxArticleid); err != nil {
            log.Fatal("Sql executed, error with result: ", err)
        }
    }
    return maxArticleid
}

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    // USE THIS VARIABLE TO LIMIT YOUR CRAWL
    crawlTime := 500
    // ANYTHING OVER 1000 WILL TAKE A LONG TIME

    c := colly.NewCollector(
        // only allow links with apnews.com as a domain
        colly.AllowedDomains("apnews.com"),
    )

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
                if h1 != "" || p != "" {
                    // reset header and para for each page
                    nextID := maxID() + 1
                    // format input
                    if len(h1) > 50 {
                        h1 = h1[0:50]
                    }
                    if len(p) > 100 {
                        p = p[0:100]
                    }
                    if len(link) > 30 {
                        insertDB(nextID, link[0:30], h1, p)
                    } else {
                        insertDB(nextID, link, h1, p)
                    }
                    header = 0
                    para = 0
                    h1 = ""
                    p = ""
                    c.Visit(e.Request.AbsoluteURL(link))
                } else {
                    c.Visit(e.Request.AbsoluteURL(link))
                }
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
}
