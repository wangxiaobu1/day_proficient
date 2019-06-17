package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_project/day_github/day_9/chatroom/common/message"
	"net"
)

//写一个函数完成登录
func login(userId int, userPwd string) (err error) {
	//1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err: ", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备消息发送消息给服务
	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes 序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal err: ", err)
		return
	}

	//5.将data赋给mes.data字段
	mes.Data = string(data)

	//6.将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json marshal err: ", err)
		return
	}

	//7.这个时候data就是我们要发的数据
	//7.1发送data长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf[4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	_, err = conn.Write(buf[0:4])
	if err != nil {
		fmt.Println("conn write header err: " , err)
		return
	}
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn write body err: " , err)
		return
	}

	//这里需要处理服务器返回的消息
	mes, err = readPkg(conn)
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}

	//将mes的data部分反序列化成LoginReMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登录成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
}
