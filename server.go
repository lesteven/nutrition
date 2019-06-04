package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
)

type Greet struct {
    Data string `json:"data"`
}



func handler(w http.ResponseWriter, r *http.Request) {
    greet := Greet{
        Data: "Welcome to the nutrition service!",
    }
    json.NewEncoder(w).Encode(greet)
}

func nutritionHandler(w http.ResponseWriter, r *http.Request) {
    nutrition := Greet{
        Data: "Nutrition Data!",
    }
    json.NewEncoder(w).Encode(nutrition)
}

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", handler)
    r.HandleFunc("/nutrition", nutritionHandler)

    log.Fatal(http.ListenAndServe(":8080", r))
}
