# go-docker

## Overview
- Print `Hello World!` with go and docker.
- Redis Access Template

## Usage

```
git clone git@github.com:urataku/go-docker.git
cd go-docker
docker-compose up -d
# Hello World!
docker exec go-docker_app_1 go run ./cmd/go-docker/main.go
# Redis Ping Pong
docker exec go-docker_app_1 go run ./cmd/go-docker/main.go
```
