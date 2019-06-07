package main

import (
    "github.com/elastic/go-elasticsearch/v8"
)

func InitElastic() *elasticsearch.Client {
    es, err := elasticsearch.NewDefaultClient()
    if err != nil {
        panic(err)
    }
    return es
}
