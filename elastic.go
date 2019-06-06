package main

import (
    "log"
    "reflect"
    "github.com/elastic/go-elasticsearch/v8"
)

func InitElastic() *elasticsearch.Client {
    es, _ := elasticsearch.NewDefaultClient()
    log.Println(reflect.TypeOf(es))
    return es
}
