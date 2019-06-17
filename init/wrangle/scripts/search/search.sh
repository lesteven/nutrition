#!/bin/bash

curl localhost:8080/elastic \
    -H 'Content-Type: application/json' \
    -d '{"search": "mochi"}' \
    -v
