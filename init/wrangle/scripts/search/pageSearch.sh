#!/bin/bash

curl "localhost:8080/elastic?size=2&offset=4"\
    -H 'Content-Type: application/json' \
    -d '{"search": "mochi"}' \
    -v
