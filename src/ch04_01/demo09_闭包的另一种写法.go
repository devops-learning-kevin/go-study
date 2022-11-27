package main

import "fmt"

func main() {
	res := func() func() int {
		i := 10
		return func() int {
			i++
			return i
		}
	}() //实现的时候调用，先调用再赋值，才可以实现闭包的效果 ()
	fmt.Printf("%T\n", res)    //func() int  //返回的时闭包函数的返回值
	fmt.Println("res:", res()) //11
	fmt.Println("res:", res()) //12
	fmt.Println("res:", res()) //13
	fmt.Printf("%T\n", res())  //int

	//如果不在实现的时候调用，结果是不一样的
	res1 := func() func() int {
		i := 10
		return func() int {
			i++
			return i
		}
	} //可以在这里调用, ()，但是现在没有调用，赋值后再调用就不一样了
	fmt.Printf("%T\n", res1)           //func() func() int 返回的是整个闭包函数
	fmt.Println("res1():", res1())     //函数的内存地址  func() int 函数的内存地址  res1(): 0x108b860
	fmt.Println("res1():", res1())     //res1(): 0x108b840
	fmt.Println("res1():", res1())     //res1(): 0x108b820
	fmt.Printf("%T\n", res1())         //func() int
	fmt.Println("res1()():", res1()()) //11
	fmt.Println("res1()():", res1()()) //11
	fmt.Println("res1()():", res1()()) //11

	fmt.Printf("%T\n", res1()()) //int

}
