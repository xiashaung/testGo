package main

import "github.com/go-redis/redis/v7"

var goRedis *redis.Client

func myredis() *redis.Client {
	if goRedis==nil {
		client :=redis.NewClient(&redis.Options{
			Addr:"127.0.0.1:6379",
			Password:"123456",
			DB:0,
		})
		goRedis = client
	}
	return goRedis
}