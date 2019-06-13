#!/bin/bash

curl localhost:9200/product/_search?pretty \
    -H 'Content-Type: application/json' \
    -d '{"query" : 
            {
               "fuzzy": { 
                    "Name": "chicken"
                }
            } 
        }'\
    -v
