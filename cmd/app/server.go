package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "nutrition/internal/route"
)

func main() {
    r := mux.NewRouter()

    // / path
    route.GreetRoute(r)

    // /nutrition path
    route.NutritionRoute(r)

    log.Fatal(http.ListenAndServe(":8080", r))
}
