package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("err dialing", err.Error())
		return
	}
	defer conn.Close()
	msg := "GET / HTTP/1.1\r\n"
	msg += "Host: www.baidu.con\r\n"
	msg += "Connection: close\r\n"
	//msg += "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36\r\n"
	msg += "\r\n\r\n"

	n, err := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed, ", err)
		return
	}
	fmt.Println("send to baidu.con bytes: ", n)
	buf := make([]byte, 4096)
	for {
		count, err := conn.Read(buf)
		fmt.Println("count: ", count, "err: ", err)
		if err != nil {
			break
		}
		fmt.Println(string(buf[0:count]))
	}
}
