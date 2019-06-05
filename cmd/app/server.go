package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "nutrition/internal/route"
)
/*
type Env struct {
    db
}
*/
func main() {
    r := mux.NewRouter()

    //  db := pgdb.InitDb()
    // / path

    r.HandleFunc("/", route.GreetHandler)
    r.HandleFunc("/nutrition", route.NutritionHandler)

    log.Fatal(http.ListenAndServe(":8080", r))
}
