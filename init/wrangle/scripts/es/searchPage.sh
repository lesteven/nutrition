#!/bin/bash

curl "localhost:9200/product/_search?size=2&from=6" \
    -H 'Content-Type: application/json' \
    -d '{"query" : 
            {
               "fuzzy": { 
                    "name": "mochi"
                }
            } 
        }'\
    -v
