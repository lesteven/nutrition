package pgdb

import (
    "fmt"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

func InitDb() *sql.DB {
    connStr := "user=steven dbname=nutrition sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func GetTime(db *sql.DB) {
    data, err := db.Query("SELECT now()")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(data)
}

