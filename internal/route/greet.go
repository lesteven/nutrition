package route

import (
    "nutrition/internal/res"
    "net/http"
    "github.com/gorilla/mux"
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

func GreetRoute(r *mux.Router) {
    r.HandleFunc("/", handler)
}
