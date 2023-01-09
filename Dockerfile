FROM golang:1.19.4-buster AS builder

ENV GOPROXY=https://goproxy.cn,direct
ENV GOPATH=/

RUN apt-get update && \
    apt-get -y install git libpq-dev postgresql-client

COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/app/main.go

ENV CGO_ENABLED 0