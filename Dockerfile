FROM golang:1.18-alpine as builder

ENV GOPATH /go/src/app
WORKDIR /go/src/app/build

COPY go.mod go.sum main.go ./

RUN go mod download

COPY cmd cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app-bin

