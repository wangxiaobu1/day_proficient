package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, ", err)
		return
	}
	defer c.Close()
	_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get book_list failed, ", err)
		return
	}
	fmt.Println(r)
}
