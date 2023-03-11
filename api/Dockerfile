# syntax=docker/dockerfile:1
FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./
RUN go build -o /api

EXPOSE 8000

ENTRYPOINT [ "/api" ]
