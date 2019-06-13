package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	conn, err := redis.Dial("tcp_client", "10.19.4.8:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Do("set", "myKey", "abc")
	if err != nil {
		log.Fatalln("set redis_connect err: ", err)
	}

	myKey, err := redis.String(conn.Do("get", "myKey"))
	if err != nil {
		log.Fatalln("get redis_connect err: ", err)
	}
	log.Println(myKey)
}
