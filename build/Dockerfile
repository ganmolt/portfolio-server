FROM golang:1.19.1-alpine

RUN apk update && apk add git

WORKDIR /go/src
RUN go mod init portfolio-server
RUN go mod tidy
