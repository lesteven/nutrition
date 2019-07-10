FROM golang:1.12 as builder


COPY . /proj

WORKDIR /proj

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o nutrition .



FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /proj/nutrition .

ENV ADDRESS="http://localhost:9200"

EXPOSE 8080

ENTRYPOINT ["./nutrition"]

