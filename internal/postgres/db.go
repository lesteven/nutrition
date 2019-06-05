package postgres

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
    fmt.Println(reflect.TypeOf(db))
    if err != nil {
        log.Fatal(err)
    }
    return db
}

func GetTime(db *sql.DB) *sql.Rows {
    data, err := db.Query("SELECT now()")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(data)
    return data
}

type FoodData struct {
    Food string `json:"food"`
    Num int `json:"num"`
}

func GetFood(db *sql.DB) FoodData {
    food := FoodData{}
    data := db.QueryRow("SELECT * FROM data where food = $1", "carrot").Scan(&food.Food, &food.Num)
    fmt.Println(data)
    fmt.Println(food)
    return food
}
