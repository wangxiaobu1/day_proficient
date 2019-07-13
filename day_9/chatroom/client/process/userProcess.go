package process

import (
	"day_daily/day_proficient/day_9/chatroom/client/utils"
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

type UserProcess struct {

}

func (this *UserProcess)Register(userId int,
	userPwd string, userName string) (err error) {
	//1. 连接到服务器
	fmt.Println("userId: ",  userId)
	fmt.Println("UserPwd: ", userPwd)
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
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserPwd = userName

	//4 将registerMes序列化
	data, err := json.Marshal(registerMes)
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
	//创建一个Transfer实例
	tf := &utils.Transfer{
		Conn : conn,
	}
	//发动data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err：", err)
	}
	mes, err = tf.ReadPkg() //mes是registerResMes
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}

	//将mes的data部分反序列化成RegisterResMes
	var registerReMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerReMes)
	if registerReMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerReMes.Error)
		os.Exit(0)
	}
	return
}

//写一个函数完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//1. 连接到服务器
	fmt.Println("userId: ",  userId)
	fmt.Println("UserPwd: ", userPwd)
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net dial err: ", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备消息发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

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

	fmt.Printf("客户端，发消息的长度=%d, 内容=%s", len(data), string(data))
	//"{\"userId\":100,\"userPwd\":\"123456\",\"userName\":\"\"}"
	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn write body err: " , err)
		return
	}
	tf := &utils.Transfer{
		Conn : conn,
	}
	//这里需要处理服务器返回的消息
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err: ", err)
		return
	}
	//将mes的data部分反序列化成LoginReMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {

		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//fmt.Println("登录成功")
		//显示当前在线用户的列表
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {
			if v == userId {
				continue
			}

			fmt.Println("用户id:\t", v)
			//完成 客户端 onlineUsers 初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUser[v] = user
		}
		fmt.Print("\n\n")
		//这里我们还需要在客户端启动一个协程
		//该协程保持和服务端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//1. 显示登录成功后的菜单
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
		//os.Exit(0)
	}
	return
}



