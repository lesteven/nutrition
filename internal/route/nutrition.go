package route

import (
    "nutrition/internal/res"
    "net/http"
    "github.com/gorilla/mux"
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


func NutritionRoute(r *mux.Router) {
    r.HandleFunc("/nutrition", nutritionHandler)
}
