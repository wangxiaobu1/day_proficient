package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func main() {

	conn, err := redis.Dial("tcp", "10.19.4.8:6379")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	_, err = conn.Do("set", "myKey", "abc")
	if err != nil {
		log.Fatalln("set redis err: ", err)
	}

	myKey, err := redis.String(conn.Do("get", "myKey"))
	if err != nil {
		log.Fatalln("get redis err: ", err)
	}
	log.Println(myKey)
}
