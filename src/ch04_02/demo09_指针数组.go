package main

import "fmt"

const COUNT int = 4

func main() {
	a := [COUNT]string{"ABC", "abc", "123", "一二三"}

	//查看数组的指针的类型和值
	fmt.Printf("%T,%v\n", &a, &a) //*[4]string,&[ABC abc 123 一二三]

	//定义指针数组
	var ptr [COUNT]*string
	fmt.Printf("%T,%v\n", ptr, ptr) //[4]*string,[<nil> <nil> <nil> <nil>]

	for i := 0; i < COUNT; i++ {
		ptr[i] = &a[i]
	}

	fmt.Println(ptr)

	for _, v := range ptr {
		fmt.Println(*v)
	}

	for i := 1; i < COUNT; i++ {
		fmt.Println(*ptr[i])
	}

}
