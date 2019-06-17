package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

//创建全局变量pool
var pool *redis.Pool

//初始化pool
func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		MaxIdle: maxIdle,
		MaxActive: maxActive,
		IdleTimeout: idleTimeout,
		Dial: func()(redis.Conn, error) {
			//return redis_conn.Dial("tcp_client", "localhosts:6379")
			return redis.Dial("tcp_client", address)
		},
	}
}
