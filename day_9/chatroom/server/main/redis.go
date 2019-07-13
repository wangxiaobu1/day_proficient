package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

//定义一个全局的pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeOut time.Duration) {
	pool = &redis.Pool{
		MaxIdle : maxIdle,//最大空闲链接数
		MaxActive : maxActive,//表示和数据库的最大链接数，0 代表没有限制
		IdleTimeout : idleTimeOut,//最大空闲时间
		Dial: func() (redis.Conn, error) {//初始化链接的代码
			return redis.Dial("tcp", address)
		},
	}
}