package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func RedisConn() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	ErrorCheck(err)
	fmt.Println(pong)
	return client
}
