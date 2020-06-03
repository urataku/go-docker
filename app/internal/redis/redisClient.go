package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var address string = "redis"
var port string = "6379"
var password string = ""
var db int = 0
var rClient *redis.Client

func Init() {
	rClient = RedisConnection()
}

func RedisConnection() *redis.Client {
	addr := address + ":" + port
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return redisClient
}

func Ping() {
	v, err := rClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(v)
}
