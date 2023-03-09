package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由  func是回调函数，用于路由的响应，
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request：包含客户端发来的数据
		fmt.Println("客户请求：", request)
		//write：通过write将数据返回给客户端
		_, _ = io.WriteString(writer, "这是/user返回的数据")
	})
	//请求指定路由，返回指定结果
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name返回的数据")
	})
	//请求指定路由，返回指定结果
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id返回的数据")
	})
	fmt.Println("http start")

	if err := http.ListenAndServe("localhost:8088", nil); err != nil {
		fmt.Println("启动服务失败", err)
		return
	}
}
