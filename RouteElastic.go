package main

import (
    "net/http"
    "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
    "encoding/json"
    "strconv"
    "log"
)

type SearchData struct {
    Search string `json:"search"`
}

func fuzzySearch(w http.ResponseWriter, es *elasticsearch.Client,
    search string, queries map[string][]string) {

    query := map[string]interface{}{
        "query": map[string]interface{}{
            "fuzzy": map[string]interface{}{
                "name": search,
            },
        },
    }
    size := 10
    offset := 0
    if len(queries["size"]) >= 1 && len(queries["offset"]) >= 1 {
        querySize, _ := strconv.Atoi(queries["size"][0])
        queryOffset, _ := strconv.Atoi(queries["offset"][0])
        if querySize >= 0 && querySize <= 20 {
            size = querySize
        }
        if queryOffset >= 0 && queryOffset <= 500 {
            offset = queryOffset
        }
    }
    res, err := es.Search(
        es.Search.WithIndex("product"),
        es.Search.WithBody(esutil.NewJSONReader(&query)),
        es.Search.WithPretty(),
        es.Search.WithSize(size),
        es.Search.WithFrom(offset),
    )

    if err != nil {
        log.Fatal(err)
        panic(err)
    }
    defer res.Body.Close()

    dataToStruct(res.Body, w)

}

func ElasticHandler (es *elasticsearch.Client) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
            case http.MethodPost:
                var search SearchData
                queries := r.URL.Query()
                err := json.NewDecoder(r.Body).Decode(&search)

                if err != nil {
                    log.Fatal(err)
                    panic(err)
                }

                fuzzySearch(w, es, search.Search, queries)
        }
    }
}

