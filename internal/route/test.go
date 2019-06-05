package route

import (
    "net/http"
    "database/sql"
    "nutrition/internal/postgres"
    "nutrition/internal/res"
)

type SqlData struct {
    Status int `json:"status"`
    Data *sql.Rows `json:"data"`
}

func TestHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := postgres.GetTime(db)
        res.SendJson(w, http.StatusOK, SqlData{http.StatusOK,data})
    }
}
