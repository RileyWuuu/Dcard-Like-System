package redis

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var db *redis.Client
var once sync.Once
var initialized bool

func GetRedis() *redis.Client {
	if !initialized {
		Initialize()
	}

	return db
}

func Initialize() {
	once.Do(func() {
		db = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
		})
		pong, err := db.Ping().Result()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(pong)
		initialized = true
	})
}
