package util

import (
	"log"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, error = rdb.Ping().Result()
	if error != nil {
		log.Println("redis 连接出错了")
	}

}
