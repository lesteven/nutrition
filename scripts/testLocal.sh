#!/bin/bash


cd "$(dirname "$0")/.."
ADDRESS=http://localhost:9200 go run *.go
