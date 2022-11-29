package main

import "fmt"

func main() {
	a := 10
	fmt.Println("调用函数之前a的值：", a)

	change1(a)
	fmt.Println("调用非指针变量的函数之后a的值：", a)

	b := &a
	change2(b)
	fmt.Println("调用指针变量的函数之后a的值：", a)

}

//参数是非指针变量
func change1(num int) {
	num = 20
}

// 参数是指针变量
func change2(num *int) {
	*num = 20
}

//总结：
//函数调用过程中，普通变量作为参数在函数内部被重新赋值修改，不会改变原实参的值，
//如果指针变量作为参数在函数内部被重新赋值修改，则实参的值同时被修改。
