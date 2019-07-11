package main

import (
    "net/http"
    "log"
)

type Greet struct {
    Status int `json:"status"`
    Data string `json:"data"`
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
    log.Printf("ADDRESS: " + address)
    greet := Greet{
        Status: http.StatusOK,
        Data: "Welcome to the nutrition service!",
    }
    SendJson(w, http.StatusOK, greet)
}

