package route

import (
    "net/http"
    "nutrition/internal/res"
    "nutrition/internal/postgres"
)

type NutData struct {
    Status int `json:"status"`
    Data string `json:"data"`
}

func NutritionHandler(w http.ResponseWriter, r *http.Request) {
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

func sqlHandler(w http.ResponseWriter, r *http.Request) {
    db := postgres.InitDb()
    postgres.GetTime(db)
    res.SendJson(w, http.StatusOK, "hello")
}

