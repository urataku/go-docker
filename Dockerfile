FROM golang:latest
RUN mkdir -p /go/src/github.com/urataku/app
WORKDIR /go/src/github.com/urataku/app
ADD . /go/src/github.com/urataku/app