package main

import (
    "net/http"
    "github.com/elastic/go-elasticsearch/v8"
    "bytes"
    "fmt"
)

type Hits struct {

}

func ElasticHandler (es *elasticsearch.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        res, err := es.Search(
            es.Search.WithIndex("customer"),
            es.Search.WithPretty(),
        )

        if err != nil {
            panic(err)
        }
        defer res.Body.Close()


        buf := new(bytes.Buffer)
        buf.ReadFrom(res.Body)
        newStr := buf.String()

        w.Header().Add("Content-Type", "application/json")
        fmt.Fprintf(w, newStr)

    }
}

