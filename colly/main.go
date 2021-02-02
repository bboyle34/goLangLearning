package main

import (
    "fmt"
    "time"
    "github.com/gocolly/colly/v2"
    "log"
    "net/http"
    "encoding/json"
    //"github.com/gofiber/fiber/v2"
)

func ping(w http.ResponseWriter, r *http.Request) {
    log.Println("Ping")
    w.Write([]byte("ping"))
}

func main() {
    fmt.Println("Hello Brendan, it is", time.Now().Format("2006-01-02 3:4:5 PM"))
    fmt.Println()

    addr := ":7171"
    //app := fiber.New()
    //defer. log.Fatal.(app.Listen(":3000"))
    //app.Use(getData)

    http.HandleFunc("/search", getData)
    http.HandleFunc("/ping", ping)

    log.Println("listening on", addr)
    log.Fatal(http.ListenAndServe(addr, nil))
}

func getData(w http.ResponseWriter, r *http.Request) {
    // verify the param "URL" exists
    URL := r.URL.Query().Get("url")
    if URL == "" {
        log.Println("missing URL argument")
        return
    }
    log.Println("visiting", URL)

    // create a new collector which will be in charge of collecting the data from HTML
    c := colly.NewCollector()

    // slices to store data
    var response []string

    // onHTML function allows the collector to use a call back function when the
    // specific HTML tag is reached
    // in this case whenever our collector finds an
    // anchor tag with href it will call the anonymous function
    // specified below which will get the info from the href and append it to our slice
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Request.AbsoluteURL(e.Attr("href"))
        if link != "" {
            log.Println(link)
            response = append(response, link)
        }
        //for _, i := range response {
        //    log.Println(response[i])
        //}
        log.Println()
        log.Println("All URLs Crawled")
        log.Println()
    })

    // command to visit the website
    c.Visit(URL)

    // parse our response slice into JSON format    
    b, err := json.Marshal(response)
    if err != nil {
        log.Println("failed to serialize response:", err)
        return
    }

    // add some header adn write the body for our endpoint
    w.Header().Add("Content-Type", "application/json")
    w.Write(b)
    //log.Println(b)
}
