package main

import (
    "reflect"
    "fmt"
    "log"
    "database/sql"
    _ "github.com/lib/pq"
)

// db is type *sql.DB
func InitDb() *sql.DB {
    connStr := "user=steven dbname=nutrition sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    log.Println(reflect.TypeOf(db))
    if err != nil {
        log.Fatal(err)
    }
    return db
}

// json (a struct) embedded within FoodData
type JsonData struct {
    Json
    FoodData `json:"data"`
}
type FoodData struct {
    Food string `json:"food"`
    Num int `json:"num"`
}

// empty interface return type to return any struct
func GetFood(db *sql.DB) interface{} {
    food := JsonData{}
    err := db.QueryRow("SELECT * FROM data where food = $1", "carrot").
           Scan(&food.Food, &food.Num)
    switch err {
        case nil:
            food.Status = 200
            return food
        case sql.ErrNoRows:
            return ErrJson{204, "There was no food"}
        default:
            panic(err)
            return ErrJson{500, "There was a server error"}
    }
}

func GetAll(db *sql.DB) interface{} {
    food := JsonData{}
    rows, err := db.Query("SELECT * FROM data LIMIT $1", 2)
    if err != nil {
        panic(err)
    }
    for rows.Next() {
        data := FoodData{}
        err = rows.Scan(&data.Food, &data.Num)
        if err != nil {
            panic(err)
        }
        fmt.Println(data)
    }
    fmt.Println("=======")
    return food
}
