package main

import "fmt"

func main() {
	res := Counter()
	fmt.Printf("%T\n", res)
	fmt.Println("res", res)
	fmt.Println("res():", res())
	fmt.Println("res():", res())
	fmt.Println("res():", res())
	//fmt.Printf("res-------------:%T,%v\n", res, res)
	//fmt.Printf("res()-----------:%T,%v\n", res(), res())
	//fmt.Printf("Counter()-------:%T,%v\n", Counter(), Counter())
	//fmt.Printf("Counter()()-----:%T,%v\n", Counter()(), Counter()())

	fmt.Printf("Counter()的类型是：%T\n", Counter())

	fmt.Println("第1次调用Counter()分配的内存地址：", Counter()) //为什么三次调用分配的内存地址都是一样的
	fmt.Println("第2次调用Counter()分配的内存地址：", Counter())
	fmt.Println("第3次调用Counter()分配的内存地址：", Counter())

	fmt.Println("----------------------------------------------------")
	//res1 := Counter1()
	//fmt.Printf("res1-------------:%T,%v\n", res1, res1)
	//fmt.Printf("res1()-----------:%T,%v\n", res1(), res1())
	//fmt.Printf("Counter1()-------:%T,%v\n", Counter1(), Counter1())
	//fmt.Printf("Counter1()()-----:%T,%v\n", Counter1()(), Counter1()())
	res1 := Counter()
	fmt.Printf("%T\n", res1)
	fmt.Println("res", res1)
	fmt.Println("res():", res1())
	fmt.Println("res():", res1())
	fmt.Println("res():", res1())

	fmt.Println("第1次调用Counter1()分配的内存地址：", Counter1()) //分什么三次抵用分配的内存地址是不同的？
	fmt.Println("第2次调用Counter1()分配的内存地址：", Counter1())
	fmt.Println("第3次调用Counter1()分配的内存地址：", Counter1())
	//函数被调用两次，他分配的内存空间是不变的。
	//结论闭包函数调用后必须赋值一个变量，然后用函数变量再去调用，计数器才会累加，否则计数器不会累加，因为不是同一个环境

}

/*
闭包函数的特性
1、保护一个在函数重复调用时不被重新初始化的变量
2、函数的返回值一定是一个函数类型，也就是返回一个内存地址
*/

func Counter() func() int {
	i := 0
	res := func() int { //把匿名函数先赋值给了一个变量，再返回这个变量。
		i++
		return i
	}
	fmt.Printf("Counter内部：%T,%v\n", res, res)
	return res
}

//闭包函数可以简写成下面这样
//下面的例子没有把匿名函数赋值给一个变量返回

func Counter1() func() int {
	i := 0
	return func() int { //直接返回匿名函数
		i++
		return i
	}
}
