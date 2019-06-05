package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "nutrition/internal/route"
    "nutrition/internal/postgres"
)

/*
type Env struct {
    db
}
*/

func main() {
    r := mux.NewRouter()

    db := postgres.InitDb()

    r.HandleFunc("/", route.GreetHandler)
    r.HandleFunc("/nutrition", route.NutritionHandler(db))
    r.HandleFunc("/test", route.TestHandler(db))

    log.Fatal(http.ListenAndServe(":8080", r))
}
