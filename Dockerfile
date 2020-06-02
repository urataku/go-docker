FROM golang:latest
RUN mkdir -p /go/src/github.com/urataku/app
WORKDIR /go/src/github.com/urataku/app
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .