package main

import "fmt"

func main() {
	var a int
	var ptr *int
	var pptr **int
	var ppptr ***int

	a = 123
	//给指针赋值
	ptr = &a
	fmt.Println("ptr:", ptr)

	//给指针的指针赋值
	pptr = &ptr
	fmt.Println("pptr:", pptr)

	ppptr = &pptr
	fmt.Println("ppptr:", ppptr)

	//获取指针对应的值
	fmt.Printf("变量 a = %d\n ", a)
	fmt.Printf("变量 *ptr = %d\n ", *ptr)
	fmt.Printf("变量 **ptr = %d\n ", **pptr)
	fmt.Printf("变量 **pptr = %d\n ", ***ppptr)

}
