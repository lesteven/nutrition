package res

import (
    "net/http"
    "encoding/json"
)

func header(w http.ResponseWriter, status int) {
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
}


func SendJson(w http.ResponseWriter, status int, data interface{}) {
    header(w, status)
    json.NewEncoder(w).Encode(data)
}

type ErrJson struct {
    Status int `json:"status"`
    Error string `json:"error"`
}

func SendError(w http.ResponseWriter, status int, errMsg string) {
    header(w, status)
    data := ErrJson {
        Status: status,
        Error: errMsg,
    }
    json.NewEncoder(w).Encode(data)
}
