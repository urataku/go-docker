# go-docker

## Overview
- Print `Hello World!` with go and docker.
- Redis Access Template
- MySQL SELECT Query Template

## Usage

```
git clone git@github.com:urataku/go-docker.git
cd go-docker
docker-compose up -d
# Hello World!
docker exec go-docker_app_1 go run ./cmd/app/main.go
# Redis Ping Pong
docker exec go-docker_app_1 go run ./cmd/redis/main.go
# MySQL 
docker exec go-docker_app_1 go run ./cmd/mysql/main.go
```
