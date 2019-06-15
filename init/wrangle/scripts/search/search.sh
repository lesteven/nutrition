#!/bin/bash

curl localhost:8080/test \
    -H 'Content-Type: application/json' \
    -d '{"search": "hello"}' \
    -v
