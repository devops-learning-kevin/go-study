package main

import "fmt"

func main() {
	// f1,f3闭包函数的非标准写法，作为返回值的函数的实现直接返回，并没有赋值给中间变量
	//直接调用函数的地址是变化的
	fmt.Printf("第1次调用f1的地址是：%v\n", f1(1))
	fmt.Printf("第2次调用f1的地址是：%v\n", f1(1))
	fmt.Printf("第3次调用f1的地址是：%v\n", f1(1))

	//赋值后再调用，函数的地址是不变的
	res3 := f3(1)
	fmt.Printf("第1次调用f3的地址是：%v\n", res3)
	fmt.Printf("第2次调用f3的地址是：%v\n", res3)
	fmt.Printf("第3次调用f3的地址是：%v\n", res3)

	//直接调用，计数器没有增长
	fmt.Printf("第1次调用f1的值是：%v\n", f1(1)())
	fmt.Printf("第2次调用f1的值是：%v\n", f1(1)())
	fmt.Printf("第3次调用f1的值是：%v\n", f1(1)())

	//赋值给变量后再调用，计数器有增长
	res := f1(1)
	fmt.Printf("第1次调用f1的值是：%v\n", res())
	fmt.Printf("第2次调用f1的值是：%v\n", res())
	fmt.Printf("第3次调用f1的值是：%v\n", res())
	fmt.Println("-----------------------------------------------")
	//闭包函数的标准写法是f2,在函数内部先把返回的函数赋值给变量，然后返回变量
	//直接调用函数的地址是变化的
	fmt.Printf("第1次调用f2的地址是：%v\n", f2(1))
	fmt.Printf("第2次调用f2的地址是：%v\n", f2(1))
	fmt.Printf("第3次调用f2的地址是：%v\n", f2(1))

	//赋值后再调用，函数的地址是不变的
	res2 := f2(1)
	fmt.Printf("第1次调用f2的地址是：%v\n", res2)
	fmt.Printf("第2次调用f2的地址是：%v\n", res2)
	fmt.Printf("第3次调用f2的地址是：%v\n", res2)

}

//注意这两种写法，完全不同
//注意闭包的返回值函数必须是：无参函数，因为你定义有参的函数，没法传参
func f1(num int) func() int {
	counter := 1
	return func() int {
		counter++
		return counter + num
	}
}

//同f1实现方式相同，只是改了个名字
func f3(num int) func() int {
	counter := 1
	return func() int {
		counter++
		return counter + num
	}
}

func f2(num int) func() int {
	counter := 1
	res := func() int {
		counter++
		return counter + num
	}
	return res
}
