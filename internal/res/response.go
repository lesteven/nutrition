package res

import (
    "net/http"
    "encoding/json"
)

func SendJson(w http.ResponseWriter, status int, data interface{}) {
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}

type ErrJson struct {
    Err string `json:"error"`
}

func SendError(w http.ResponseWriter, status int, errMsg string) {
    w.WriteHeader(status)
    w.Header().Add("Content-Type", "application/json")
    data := ErrJson {
        Err: errMsg,
    }
    json.NewEncoder(w).Encode(data)
}
