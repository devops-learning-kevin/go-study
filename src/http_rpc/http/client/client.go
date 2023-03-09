package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}
	//func (c *Client) Get(url string) (resp *Response, err error)
	//get到需要测试的url
	resp, err := client.Get("http://localhost:8088")
	if err != nil {
		fmt.Println("请求失败", err)
		return
	}

	//get到header的指定key
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	sr := resp.Header.Get("Server")
	//输出打印结果
	//fmt.Println(ct, date, sr)
	fmt.Println("ct=====", ct)
	fmt.Println("date=====", date)
	fmt.Println("sr=====", sr)

	//获取请求的url，状态码，状态
	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status
	//输出打印结果
	//fmt.Println(url, code, status)
	fmt.Println("url=====", url)
	fmt.Println("code=====", code)
	fmt.Println("status=====", status)

	//拿到Body体内的所有内容
	body := resp.Body
	//输出打印body
	fmt.Println("body =====", body)
	//接收body转换为string类型
	readAll, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("请求body失败", err)
		return
	}
	//输出结果
	fmt.Println("readAll======", string(readAll))
}
