# syntax=docker/dockerfile:1

FROM golang:1.21-alpine

WORKDIR /app

COPY . ./

RUN go build -o ./docker-gs-ping

CMD ["./docker-gs-ping"]
