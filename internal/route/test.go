package route

import (
    "net/http"
    "database/sql"
    "nutrition/internal/postgres"
    "nutrition/internal/res"
)


func TestHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        data := postgres.GetFood(db)
        res.SendJson(w, http.StatusOK, data)
    }
}
