package main

import (
	"fmt"
)

/*
	值传递：是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中对参数进行修改，不会影响到实际参数。
	引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数的修改，将会影响到实际参数。
	严格来说Go语言只有值传递一种传参方式，Go语言是没有引用传递的。
	Go语言可以借助传指针来实现引用传递的效果。函数参数使用指针参数，传参时其实是在拷贝一份指针参数，也就是拷贝一份变量地址。

	值传递和引用传递的注意细节【重要】
	1、Go语言中所有的传参都是值传递(传值)，都是一个副本，一个拷贝。
	。拷贝的内容有时候是非引用类型(int、string、 bool、数组、struct属于非引用类型)，这样就在函数中就无
	法修改原内容数据;
	。有的是引用类型(指针、slice、map、chan属属于引用类型)，这样就可以修改原内容数据。
	2、是否可以修改原内容数据，和传值、传引用没有必然的关系。在C++中，传引用肯定是可以修改原内容数据的，在Go语言里，虽然只有传值，但是我们也可以修改原内容数据，因为参数可以是引用类型。
	3、传引用和引用类型是两个概念。虽然Go语言只有传值直一种方式，但是可以通过传引用类型变量达到跟传引用一样的效果。
*/

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Printf("1、变量a的内存地址：%p,值为：%v\n\n", &a, a) //[1,2,3,4]

	//传值,特别注意，切片slice是引用类型，传值实际传递的也是地址
	changeSliceVal(a)
	fmt.Printf("2、changeSliceVal调用后：变量a的内存地址：%p,值为：%v\n\n", &a, a) // [1,2,3,4]X [90,2,3,4]

	//传引用
	changeSlicePtr(&a)
	fmt.Printf("3、changeSlicePtr调用后：变量a的内存地址：%p,值为：%v\n\n", &a, a) // [1,250,3,4]X [90,250,3,4]
}

func changeSliceVal(a []int) {
	fmt.Printf("----------changeSliceVal函数内：参数a的内存地址：%p,值为%v\n", &a, a) // &[90,2,3,4]
	a[0] = 90
}

func changeSlicePtr(a *[]int) {
	fmt.Printf("----------changeSlicePtr函数内：参数a的内存地址：%p,值为%v\n", &a, a) //地址
	(*a)[1] = 250

}
