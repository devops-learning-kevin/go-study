package main

import "fmt"

/*
	值传递：是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中对参数进行修改，不会影响到实际参数。
	引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数的修改，将会影响到实际参数。
	严格来说Go语言只有值传递一种传参方式，Go语言是没有引用传递的。
	Go语言可以借助传指针来实现引用传递的效果。函数参数使用指针参数，传参时其实是在拷贝一份指针参数，也就是拷贝一份变量地址。
*/

func main() {
	a := 10
	fmt.Printf("1、变量a的内存地址：%p,值为：%v\n\n", &a, a) //10

	//传值
	changeIntVal(a)
	fmt.Printf("2、changeIntVal调用后：变量a的内存地址：%p,值为：%v\n\n", &a, a) //10

	//传引用
	changeIntPtr(&a)
	fmt.Printf("3、changeIntPtr调用后：变量a的内存地址：%p,值为：%v\n\n", &a, a) //50
}

func changeIntVal(a int) {
	fmt.Printf("----------changeIntVal函数内：参数a的内存地址：%p,值为%v\n", &a, a)
	a = 90
}

func changeIntPtr(a *int) {
	fmt.Printf("----------changeIntPtr函数内：参数a的内存地址：%p,值为%v\n", &a, a)
	*a = 50

}
