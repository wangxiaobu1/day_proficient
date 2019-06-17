package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed, ", err)
	}
	defer c.Close()
	_, err = c.Do("HSet", "book", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	r, err := redis.Int(c.Do("HGet", "book", "abc"))
	if err != nil {
		fmt.Println("get abc failed, ", err)
		return
	}
	fmt.Println(r)
}
