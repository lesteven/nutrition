package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    db := InitDb()
    InitElastic()

    r.HandleFunc("/", GreetHandler)
    r.HandleFunc("/nutrition", NutritionHandler(db))
    r.HandleFunc("/test", TestHandler(db))

    log.Fatal(http.ListenAndServe(":8080", r))
}
