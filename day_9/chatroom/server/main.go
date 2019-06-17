package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"encoding/binary"
	"go_project/day_github/day_9/chatroom/common/message"
	"io"
	"net"
)

func readPkg(conn net.Conn,) (mes message.Message, err error) {
	buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	_, err = conn.Read(buf[0:4])
	if err != nil {
		err = errors.New("read pkg header error")
		return
	}
	//根据buf[0:4]转成一个int32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(buf[0:4])
	//根据pkgLen读取消息内容
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	//吧pkgLen 反序列化成 message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmarshal err= ", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {
	//发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(bytes) fail ", err)
		return
	}
	//发送data本身
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write fail ", err)
		return
	}
	return
}

//编写一个函数serverProcessLogin函数，专门处理登录信息
func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
	//1.从message取出mes.data，并反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json unmarshal fail err: ", err)
		return
	}

	//2.先声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//3.再声明LoginResMes
	var loginResMes message.LoginResMes

	//如果用户id 是100，密码是123456 是合法，否则不合法
	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
		//合法
		loginResMes.Code = 200

	} else {
		//不合法
		loginResMes.Code = 500
		loginResMes.Error = "该用户不存在，请注册再使用..."
	}

	//4.将logResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}

	//5.将data赋值给resMes
	resMes.Data = string(data)

	//6.将resMes 序列化，并发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json marshal fail ", err)
		return
	}

	//7.发送data，将其封装到write函数里
	err = writePkg(conn, data)
	return
}

//编写一个serverProcessMes函数
//功能：根据消息客户端发送消息种类不同，决定调用那个函数来处理
func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		//处理登录
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

//处理客户端通信
func process(conn net.Conn) {
	//这里需要延时关闭
	defer conn.Close()
	//循环的客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个readpkg函数，返回message，err
		mes, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出...")
				return
			} else {
				fmt.Println("readpkg err: ", err)
				return
			}
		}
		fmt.Println(mes)
		err = serverProcessMes(conn, &mes)
		if err != nil {
			return
		}
	}
}

func main() {

}
