package route

import (
    "net/http"
    "github.com/gorilla/mux"
    "nutrition/internal/res"
    "nutrition/internal/pgdb"
)

type NutData struct {
    Status int `json:"status"`
    Data string `json:"data"`
}

func nutritionHandler(w http.ResponseWriter, r *http.Request) {
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
    db := pgdb.InitDb()
    pgdb.GetTime(db)
    res.SendJson(w, http.StatusOK, "hello")
}

func NutritionRoute(r *mux.Router) {
    r.HandleFunc("/nutrition", nutritionHandler)
    r.HandleFunc("/sql", sqlHandler)
}
