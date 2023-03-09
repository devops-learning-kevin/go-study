package main

import (
	"fmt"
	"net"
)

func handle(conn net.Conn) { //处理连接方法
	defer conn.Close() //关闭连接
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf) //读取客户端数据
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("接收数据长度： %d\n接收输入内容:%s\n", n, string(buf[0:n]))
		msg := []byte("hello,world\n")
		conn.Write(msg) //发送数据
	}
}
func main() {
	fmt.Println("start server....")
	listen, err := net.Listen("tcp", "0.0.0.0:3000") //创建监听
	if err != nil {
		fmt.Println("listen failed! msg :", err)
		return
	}
	for {
		conn, errs := listen.Accept() //接受客户端连接
		if errs != nil {
			fmt.Println("accept failed")
			continue
		}
		go handle(conn) //处理连接
	}
}
