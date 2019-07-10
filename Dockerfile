FROM golang:1.12

LABEL version="1.0"
LABEL decsription="nutrition server"

COPY nutrition nutrition

EXPOSE 8080

ENV ADDRESS="http://localhost:9200"

ENTRYPOINT ["./nutrition"]

