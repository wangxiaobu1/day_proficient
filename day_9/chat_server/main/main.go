package main

import "time"

func main() {
	InitRedis("localhost:6379", 16, 1024, 300*time.Second)
}
