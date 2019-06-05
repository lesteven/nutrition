package route

import (
    "net/http"
    "database/sql"
    "nutrition/internal/res"
)

type NutData struct {
    Status int `json:"status"`
    Data string `json:"data"`
}


func NutritionHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodGet:
                nutrition := NutData{
                    Status: http.StatusOK,
                    Data: "Nutrition Data!",
                }
                res.SendJson(w, http.StatusOK, nutrition)
            case http.MethodPost:
                res.SendError(w, http.StatusBadRequest, "cant post!")
        }
    }
}


