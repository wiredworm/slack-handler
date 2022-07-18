FROM golang:1.18-alpine as builder

ENV GOPATH /go/src/app
WORKDIR /go/src/app/build

COPY go.mod go.sum main.go ./
