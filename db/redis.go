package db

import (
	"fmt"
	"github.com/go-redis/redis"
)

const (
	RedisAddr = "localhost:6379"
	Password  = ""
	DB        = 0
)

var RedisClient *redis.Client

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: Password,
		DB:       DB,
	})

	pong := client.Ping().Val()
	fmt.Println(pong)

	return client
}

func init() {
	RedisClient = createClient()
}
