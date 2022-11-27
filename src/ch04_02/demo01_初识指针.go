package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("变量的地址：%x\n", &a)
	//&a++  //指针类型不能参与运算

	b := []int{1, 3, 5, 7}
	fmt.Printf("切片变量的地址：%x\n", &b)
}
