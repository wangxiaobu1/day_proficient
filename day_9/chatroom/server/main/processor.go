package main

import (
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"day_daily/day_proficient/day_9/chatroom/server/process"
	"day_daily/day_proficient/day_9/chatroom/server/utils"
	"fmt"
	"io"
	"net"
)

//创建一个processor的结构体
type Processor struct {
	Conn net.Conn
}

//编写一个serverProcessMes函数
//功能：根据消息客户端发送消息种类不同，决定调用那个函数来处理
func (this *Processor) serverProcessMes(mes *message.Message) (err error) {
	//fmt.Println("serverProcessMes")
	//看看是否能接收到客户端发送的群发消息
	fmt.Println("mes= ", mes)

	switch mes.Type {
	case message.LoginMesType:
		up := &process2.UserProcess{
			Conn : this.Conn,
		}
		//处理登录
		err = up.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
		up := &process2.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessRegister(mes)
	case message.SmsMesType:
		smsProcess := &process2.SmsProcess{}
		smsProcess.SendGroupMes(mes)
	default:
		fmt.Println("消息类型不存在，无法处理...")
	}
	return
}

func (this *Processor) process2() (err error) {
	//循环的客户端发送的消息
	for {
		//这里我们将读取数据包，直接封装成一个readpkg函数，返回message，err
		tf := &utils.Transfer{
			Conn : this.Conn,
		}
		mes, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出，服务器也退出...")
				return err
			} else {
				fmt.Println("readpkg err: ", err)
				return err
			}
		}
		fmt.Println("mes: ", mes)
		err = this.serverProcessMes(&mes)
		if err != nil {
			return err
		}
	}
}
