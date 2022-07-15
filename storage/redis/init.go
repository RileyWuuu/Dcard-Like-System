package redis

import (
	"dcard/config"
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
			Addr:     config.Conf.Redis.Addr,
			Password: config.Conf.Redis.Password,
			DB:       0,
		})
		_, err := db.Ping().Result()
		if err != nil {
			fmt.Println(err)
		}
		initialized = true
	})
}
