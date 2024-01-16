FROM golang:1.19.1-alpine

RUN apk update && apk add git

# air
RUN go install github.com/cosmtrek/air@v1.29.0

WORKDIR /go/src
# module
RUN go mod init portfolio-server
RUN go mod tidy
