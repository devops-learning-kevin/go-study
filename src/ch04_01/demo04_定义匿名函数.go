package main

import (
	"fmt"
	"math"
)

func main() {
	//1、定义时调用无返回值的匿名函数
	func(data int) {
		fmt.Println("你的成绩是：", data)
	}(96)

	//2、定义时调用有返回值的匿名函数
	result := func(data float64) float64 {
		return math.Sqrt(data)
	}(250)
	fmt.Println("平方根：", result)

	//3、将匿名函数赋值给变量，需要时再调用
	myfunc := func(data string) string {
		return data
	}
	fmt.Println(myfunc("hello golang"))

}
