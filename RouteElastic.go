package main

import (
    "net/http"
    "github.com/elastic/go-elasticsearch/v8"
//    "bytes"
   "fmt"
    "encoding/json"
    "io/ioutil"
    "reflect"
)

type Results struct {

}

type Hits struct {
    Took int `json:"took"`
    TimedOut bool `json:"timed_out"`
    Shards struct {
        Total int `json:"total"`
        Success int `json:"successful"`
        skipped int `json:"skipped"`
        failed int `json:"failed""`
    } `json:"_shards"` 
    hits struct {
        Total struct {
            Value int `json:"value"`
            Relation string `json:relation"`
        } `json:"total"`
        Score float32 `json:"max_score"`
        Hits []interface{} `json: "hits"`
    } `json: "hits"`
}

type Data map[string]map[string]interface{}

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

/*
        buf := new(bytes.Buffer)
        buf.ReadFrom(res.Body)
        newStr := buf.String()

        */

        var h Data
        body, _ := ioutil.ReadAll(res.Body)
        err = json.Unmarshal(body, &h)
        fmt.Println(h["hits"]["hits"])
        fmt.Println(reflect.TypeOf(h["hits"]["hits"]))

        SendJson(w, 200, h["hits"]["hits"])
    /*
        var h 
        w.Header().Add("Content-Type", "application/json")
        json.NewEncoder(w).Encode(h)
        fmt.Println(h)
        //fmt.Fprintf(w, newStr)
*/
    }
}

