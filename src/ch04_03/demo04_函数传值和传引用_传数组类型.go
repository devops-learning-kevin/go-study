package main

import (
	"fmt"
)

/*
	值传递：是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中对参数进行修改，不会影响到实际参数。
	引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数的修改，将会影响到实际参数。
	严格来说Go语言只有值传递一种传参方式，Go语言是没有引用传递的。
	Go语言可以借助传指针来实现引用传递的效果。函数参数使用指针参数，传参时其实是在拷贝一份指针参数，也就是拷贝一份变量地址。
*/

func main() {
	a := [4]int{1, 2, 3, 4}
	fmt.Printf("1、变量a的内存地址：%p,值为：%v\n\n", &a, a) //[1,2,3,4]

	//传值
	changeArrVal(a)
	fmt.Printf("2、changeArrVal调用后：变量a的内存地址：%p,值为：%v\n\n", &a, a) //[1,2,3,4]

	//传引用
	changeArrPtr(&a)
	fmt.Printf("3、changeArrPtr调用后：变量a的内存地址：%p,值为：%v\n\n", &a, a) ////[1,250,3,4]
}

func changeArrVal(a [4]int) {
	fmt.Printf("----------changeArrVal函数内：参数a的内存地址：%p,值为%v\n", &a, a) // &[1,2,3,4]
	a[0] = 90
}

func changeArrPtr(a *[4]int) {
	fmt.Printf("----------changeArrPtr函数内：参数a的内存地址：%p,值为%v\n", &a, a) //地址
	(*a)[1] = 250

}
