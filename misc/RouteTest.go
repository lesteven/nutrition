package main

import (
    "net/http"
    "database/sql"
    "fmt"
    "encoding/json"
    //"io/ioutil"
)


func TestHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodGet:
                data := GetFood(db)
                SendJson(w, http.StatusOK, data)
            case http.MethodPost:
                var search SearchData

                err := json.NewDecoder(r.Body).Decode(&search)

                if err != nil {
                    fmt.Println(err)
                    fmt.Println("error with decoder")
                    panic(err)
                }

                fmt.Println(search.Search)
                fmt.Fprintf(w, "hello post!") 

                /*
                fmt.Println(r.Body)
                data, err := ioutil.ReadAll(r.Body)
                if err != nil {
                    panic(err)
                }
                fmt.Println(data)
                fmt.Println(string(data))
                */
        }
    }
}
