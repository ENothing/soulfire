package db

import (
	"github.com/go-redis/redis/v7"
	"time"
)

var RedisDb *redis.Client


func RedisInit()  {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:"127.0.0.1:6379",
		//DB:1,
		Password:"GUhcF7bSK?u@Rtp",
		//Password:"",
		MinIdleConns:1,
		MaxRetries:3,
		IdleTimeout: 10*time.Second,
		PoolSize:10,

	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		panic(err)
	}
}
