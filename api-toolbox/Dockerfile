# alpine:3.13.6
FROM alpine@sha256:e15947432b813e8ffa90165da919953e2ce850bef511a0ad1287d7cb86de84b5

RUN apk update && apk add --no-cache git

RUN go install github.com/OJ/gobuster/v3@latest
RUN go get -u -v github.com/ffuf/ffuf
