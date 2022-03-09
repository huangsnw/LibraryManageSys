package util

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := Redis.Ping(ctx).Result()
	log.Println("Redis Connect Sucess")
	if err != nil {
		log.Println("Redis Connect Errror")
	}
}

func Set(key string, value string) {
	err := Redis.Set(context.Background(), key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

/**
 * @description: 读取Redis
 * @Accept:
 * @param {string} key
 * @return {*}
 * @Router:
 */
func Get(key string) string {
	value, err := Redis.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	return value
}
