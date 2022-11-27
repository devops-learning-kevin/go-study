package main

import "fmt"

/*
闭包就是能够读取其他函数内部变量的函数。
例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。
在本质上，闭包是将函数内部和函数外部连接起来的桥梁。
*/

func main() {
	res := adder()
	fmt.Printf("%T \n", res)

	for i := 0; i < 5; i++ {
		fmt.Printf("%d \t", i)
		//输出计时器累加过程
		fmt.Printf("%d \t", adder()(i)) //直接调用闭包函数达不到效果
		fmt.Println(res(i))             //必须要把闭包函数赋值一个引用，建立一个引用环境。
	}
}

//实现计数器的闭包函数
//返回值是函数的函数是闭包函数
func adder() func(int) int {
	sum := 0
	return func(num int) int {
		sum += num
		return sum
	}

}
