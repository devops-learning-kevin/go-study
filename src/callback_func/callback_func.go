package main

import (
	"fmt"
	"strconv"
)

//一个记录日志的类型：func(string)
type saveLog func(msg string)

//将字符串转换为int64,如果转换失败调用saveLog
func stringToInt(s string, log saveLog) int64 {

	if value, err := strconv.ParseInt(s, 16, 32); err != nil {
		log(err.Error())
		return 0
	} else {
		return value
	}
}

//记录日志消息的具体实现
func myLog(msg string) {
	fmt.Println("Find Error:", msg)
}

func main() {
	v32 := "-354634382"
	fmt.Println(stringToInt("123", myLog)) //转换时将调用mylog记录日志
	fmt.Println(stringToInt("s", myLog))
	fmt.Println(stringToInt(v32, myLog))
	//fmt.Println(stringToInt("s", myLog))

	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
}
