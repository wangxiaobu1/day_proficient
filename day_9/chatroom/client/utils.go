package main

import (
	"encoding/json"
	"fmt"
	"net"
	"go_project/day_github/day_9/chatroom/common/message"
	"errors"
	"encoding/binary"
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
