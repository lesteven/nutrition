package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

//    db := InitDb()
    es := InitElastic()

    r.HandleFunc("/", GreetHandler)
    r.HandleFunc("/elastic", ElasticHandler(es))

    log.Fatal(http.ListenAndServe(":8080", r))
}
