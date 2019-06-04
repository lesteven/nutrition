package route

import (
    "nutrition/internal/res"
    "net/http"
    "github.com/gorilla/mux"
)

type NutData struct {
    Data string `json:"data"`
}

func nutritionHandler(w http.ResponseWriter, r *http.Request) {
    nutrition := NutData{
        Data: "Nutrition Data!",
    }
    res.SendSuccess(w, nutrition)
}

func NutritionRoute(r *mux.Router) {
    r.HandleFunc("/nutrition", nutritionHandler)
}
