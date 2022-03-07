/*
 * @title: Redis 配置文件
 * @Date: 2022-02-21 10:30:02
 * @version: 1.0
 * @author: huang sn
 * @description: 这里写描述信息
 * @FilePath: /LibraryManageSys/util/Redis.go
 */
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

/**
 * @description: 写入Redis
 * @Accept:
 * @param {string} key
 * @param {string} value
 * @return {*}
 * @Router:
 */
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
