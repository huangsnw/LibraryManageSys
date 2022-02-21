package util

import "github.com/go-redis/redis"

var Redis *redis.Client

func InitRedis() {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}
