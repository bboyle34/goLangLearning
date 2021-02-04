package main

import (
    "database/sql"
    "log"

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

    // delete existing tables
    if _, err := db.Exec(
        `delete from article;`); err != nil {
        log.Fatal("Delete statement did not work: ", err)
    }

    // insert two new tables
    if _, err := db.Exec(
        `insert into article (articleid, link, header, paragraph) 
        values (1, 'https://apnews.com', 'First Header', 'First Paragraph');`); err != nil {
        log.Fatal("Insert statement did not work: ", err)
    }
    if _, err := db.Exec(
        `insert into article (articleid, link, header, paragraph) 
        values (2, 'https://apnews.com/article/example', 'Second Header', 'Second Paragraph');`); err != nil {
        log.Fatal("Insert statement2 did not work: ", err)
    }

}
