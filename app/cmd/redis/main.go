package main

import (
	"github.com/urataku/go-docker/internal/redis"
)

func main() {
	redis.Init()
	redis.Ping()
}
