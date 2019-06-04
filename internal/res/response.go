package res

import (
    "net/http"
    "encoding/json"
)

func SendSuccess(w http.ResponseWriter, data interface{}) {
    w.Header().Add("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}
