package utils

import (
	"day_daily/day_proficient/day_9/chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取服务端发送的数据...")
	_, err = this.Conn.Read(this.Buf[0:4])
	if err != nil {
		err = errors.New("read pkg header error")
		return
	}
	//根据buf[0:4]转成一个int32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[0:4])
	//根据pkgLen读取消息内容
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body error")
		return
	}
	//把pkgLen 反序列化成 message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json unmarshal err= ", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	fmt.Println("发送给服务端包")
	//发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buf[0:4])
	if n != 4 || err != nil {
		fmt.Println("conn.write(bytes) fail ", err)
		return
	}
	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.write fail ", err)
		return
	}
	return
}
