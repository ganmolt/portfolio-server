FROM golang:1.21.4

WORKDIR /go/src
# module
RUN go mod init portfolio-server
# air
RUN go install github.com/cosmtrek/air@latest
# sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest
RUN go mod tidy
