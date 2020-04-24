package db

import (
	"github.com/go-redis/redis/v7"
)

var RedisDb *redis.Client


func RedisInit()  {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis地址
		Password: "GUhcF7bSK?u@Rtp",  // Redis账号
		DB:       1,   // Redis库
		//PoolSize: 10,  // Redis连接池大小
		//MaxRetries: 3,              // 最大重试次数
		//IdleTimeout: 10*time.Second,            // 空闲链接超时时间
	})
	_, err := RedisDb.Ping().Result()
	if err != nil {
		panic(err)
	}
}
