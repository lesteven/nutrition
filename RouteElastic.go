package main

import (
    "net/http"
    "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
    "fmt"
    "encoding/json"
)

type SearchData struct {
    Search string `json:"search"`
}

func searchES(w http.ResponseWriter, es *elasticsearch.Client, search string) {
    res, err := es.Search(
        es.Search.WithIndex("product"),
        es.Search.WithPretty(),
    )

    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    dataToMap(res.Body, w)
}

func fuzzySearch(w http.ResponseWriter, es *elasticsearch.Client, search string) {
    query := map[string]interface{}{
        "query": map[string]interface{}{
            "fuzzy": map[string]interface{}{
                "Name": search,
            },
        },
    }
    res, err := es.Search(
        es.Search.WithIndex("product"),
        es.Search.WithBody(esutil.NewJSONReader(&query)),
        es.Search.WithPretty(),
    )

    if err != nil {
        panic(err)
    }
    defer res.Body.Close()

    dataToMap(res.Body, w)

}

func ElasticHandler (es *elasticsearch.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodPost:
                var search SearchData

                err := json.NewDecoder(r.Body).Decode(&search)

                if err != nil {
                    fmt.Println(err)
                    fmt.Println("error with decoder")
                    panic(err)
                }

                fmt.Println(search.Search)
                fuzzySearch(w, es, search.Search)
        }
    }
}

