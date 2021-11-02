# golang:alpine3.14
FROM golang@sha256:5ce2785c82a96349131913879a603fc44d9c10d438f61bba84ee6a1ef03f6c6f

RUN apk update && apk add --no-cache git

RUN go install github.com/OJ/gobuster/v3@latest
RUN go get -u -v github.com/ffuf/ffuf
