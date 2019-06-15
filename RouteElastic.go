package main

import (
    "net/http"
    "github.com/elastic/go-elasticsearch/v8"
    "bytes"
    "fmt"
    "encoding/json"
    "io"
    "io/ioutil"
)

type EsData struct {
    Took int `json:"took"`
    TimedOut bool `json:"timed_out"`
    Shards struct {
        Total int `json:"total"`
        Success int `json:"successful"`
        Skipped int `json:"skipped"`
        Failed int `json:"failed""`
    } `json:"_shards"` 
    Hits struct {
        Total struct {
            Value int `json:"value"`
            Relation string `json:"relation"`
        } `json:"total"`
        Score float32 `json:"max_score"`
        Hits []map[string]interface{} `json:"hits"`
    } `json:"hits"`
}

type Data map[string]map[string][]interface{}

func dataToString(body io.ReadCloser, w http.ResponseWriter) {
    buf := new(bytes.Buffer)
    buf.ReadFrom(body)
    newStr := buf.String()
    w.Header().Add("Content-Type", "application/json")
    fmt.Fprintf(w, newStr)
}

func dataToMap(body io.ReadCloser, w http.ResponseWriter) {
    var h Data
    data, _ := ioutil.ReadAll(body)
    json.Unmarshal(data, &h)
    /*
    for _, each := range h["hits"]["hits"] {
        fmt.Println(each)
    }
    */

    SendJson(w, 200, h["hits"]["hits"])
}

func dataToStruct(body io.ReadCloser, w http.ResponseWriter) {
    h := EsData{}
    w.Header().Add("Content-Type", "application/json")
    json.NewDecoder(body).Decode(&h)
    /*
    for _, each := range h.Hits.Hits {
        fmt.Println(each["_source"])
    }
    */
    json.NewEncoder(w).Encode(h)
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

        dataToMap(res.Body, w)

    }
}

