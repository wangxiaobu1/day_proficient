package main

import (
	"log"
	"time"
	"go_project/day_redis/oneRedis/model"
)

func init() {
	initPool("localhosts:6379", 8, 0, 300*time.Second)
	initUserDao()
}

func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	user, err := model.MyUserDao.Login(100, "123456")
	if err != nil {
		log.Fatalln("出错误")
	}
	log.Println(user)
}
