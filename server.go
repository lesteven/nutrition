package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    db := InitDb()
    es := InitElastic()
    res, err := es.Search(
        es.Search.WithIndex("customer"),
        es.Search.WithPretty(),
    )
    if err != nil {
        panic(err)
    }
    log.Println(res)
    r.HandleFunc("/", GreetHandler)
    r.HandleFunc("/elastic", ElasticHandler)
    r.HandleFunc("/nutrition", NutritionHandler(db))
    r.HandleFunc("/test", TestHandler(db))

    log.Fatal(http.ListenAndServe(":8080", r))
}
