package main

import "fmt"

func main() {
	fmt.Println(GetMultiple(10))
	fmt.Println(factorial(10))

}

//使用递归要注意栈溢出
//用递归方法实现阶乘
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

//通过循环方式实现阶乘
func GetMultiple(num int) (result int) {
	result = 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return

}
