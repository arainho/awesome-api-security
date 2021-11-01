# golang:alpine3.14
FROM golang@sha256:d09ad73822f06b4aa3c1f25842299e0f1c6e7937b2785cb6e9eadeea44b18d76

RUN apk update && apk add --no-cache git

RUN go install github.com/OJ/gobuster/v3@latest
RUN go get -u -v github.com/ffuf/ffuf
