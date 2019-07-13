package main

import (
	"day_daily/day_proficient/day_9/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

//处理客户端通信
func process(conn net.Conn) {
	//这里需要延时关闭
	defer conn.Close()
	processor := &Processor{
		Conn : conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器的通讯协程错误err: ", err)
		return
	}
}

//编写一个函数，完成对userDao的初始化任务
func initUserDao() {
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	//当服务器启动时，我们就初始化我们的redis链接池
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
	//提示信息
	fmt.Println("服务器在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net listen err: ", err)
		return
	}
	//一旦监听成功就等待客户端连接服务器
	for {
		fmt.Println("等待客户端来链接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen Accept err: ", err)
		}
		go process(conn)
	}
}
