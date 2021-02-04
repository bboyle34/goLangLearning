package main

import (
    "database/sql"
    "log"
    "fmt"

    _ "github.com/lib/pq"
)

type Article struct {
    articleid int
    link string
    header string
    paragraph string
}

func main() {
    // connect to the database
    db, err := sql.Open("postgres", "postgresql://root@localhost:26257/apnews_db?sslmode=disable")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    /*
    // insert a row into the article table
    if _, err := db.Exec(
        `insert into article (articleid, link, header, paragraph) 
        values (2, 'http://article.com', 'second header', 'second paragraph')`); err != nil {
        log.Fatal(err)
    }
    */

    // get highest articleid from article table
    rows, err := db.Query("select max(articleid) from article;")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    for rows.Next() {
        var maxArticleid int
        if err := rows.Scan(&maxArticleid); err != nil {
            log.Fatal(err)
        }
        fmt.Println("The max articleid is:", maxArticleid)
    }
}
