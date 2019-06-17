package main

import (
    "github.com/elastic/go-elasticsearch/v8"
    "net/http"
    "fmt"
    "io"
    "io/ioutil"
    "bytes"
    "encoding/json"
)

func InitElastic() *elasticsearch.Client {
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        panic(err)
    }
    return es
}


// get json data and turn it into a string
func dataToString(body io.ReadCloser, w http.ResponseWriter) {
    buf := new(bytes.Buffer)
    buf.ReadFrom(body)
    newStr := buf.String()
    w.Header().Add("Content-Type", "application/json")
    fmt.Fprintf(w, newStr)
}

// { hits: { hits: []interface{} } }
type Data map[string]map[string][]interface{}

// get json data and unmarshall into map 
func dataToMap(body io.ReadCloser, w http.ResponseWriter) {
    var h Data
    data, _ := ioutil.ReadAll(body)
    json.Unmarshal(data, &h)
    SendJson(w, 200, h["hits"]["hits"])
}

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

// get json data and decode into EsData
func dataToStruct(body io.ReadCloser, w http.ResponseWriter) {
    h := EsData{}
    w.Header().Add("Content-Type", "application/json")
    json.NewDecoder(body).Decode(&h)
    json.NewEncoder(w).Encode(h.Hits)
}
