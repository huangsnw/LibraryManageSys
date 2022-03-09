package util

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.poolSize"),
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

func Get(key string) string {
	value, err := Redis.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	return value
}
