package main

import "fmt"

func main() {
	a := 10
	//b := &a  两种写法均可
	var b *int = &a

	fmt.Printf("b的类型：%T，b的值：%v\n", b, b)
	fmt.Println("a的地址：", b)
	fmt.Println("*b的值", *b) //10

	//当我们操作一个指针的时候，实际上操作的时这个指针所指向的变量
	*b++                    // 相当于 a++
	fmt.Println("a的值", a) //11

}
