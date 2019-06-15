package main

import (
    "net/http"
    "database/sql"
)


func TestHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := GetFood(db)
        SendJson(w, http.StatusOK, data)
    }
}
