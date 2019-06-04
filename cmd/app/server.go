package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "nutrition/internal/res"
)

type Greet struct {
    Data string `json:"data"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    greet := Greet{
        Data: "Welcome to the nutrition service!",
    }
    res.SendSuccess(w, greet)
}

func nutritionHandler(w http.ResponseWriter, r *http.Request) {
    nutrition := Greet{
        Data: "Nutrition Data!",
    }
    res.SendSuccess(w, nutrition)
}


func main() {
    r := mux.NewRouter()

    r.HandleFunc("/", handler)
    r.HandleFunc("/nutrition", nutritionHandler)

    log.Fatal(http.ListenAndServe(":8080", r))
}
