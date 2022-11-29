package main

import "fmt"

func main() {
	//定义两个局部变量
	a, b := 100, 200

	//使用返回值的方法实现两个变量值的交换
	a, b = swap0(a, b)
	fmt.Println(a, b)

	//使用指针作为参数的方法来实现两个变量值的交换
	swap(&a, &b)
	fmt.Println(a, b)

}

//传统方法（使用返回值的方法）交换两个变量的值
func swap0(x, y int) (int, int) {
	return y, x
}

//使用指针来交换两个变量的值
func swap(x, y *int) {
	*x, *y = *y, *x
}
