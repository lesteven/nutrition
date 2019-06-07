package main

import (
    "net/http"
    "fmt"
)


func ElasticHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reached elastic!")
    resp, err := http.Get("http://localhost:9200/_search?pretty")
    if err != nil {
        panic(err)
    }
    SendJson(w, http.StatusOK, resp)
}

