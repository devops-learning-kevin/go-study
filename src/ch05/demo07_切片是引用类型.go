package main

import "fmt"

func main() {
	//数组是值类型
	a := [4]int{1, 2, 3, 4}

	//切片是引用类型
	b := []int{100, 200, 300}

	c := a
	d := b

	c[1] = 200
	d[0] = 1

	fmt.Println("a=", a, "c=", c) //a=[1 2 3 4 ]  c=[1 200 3 4]

	fmt.Println("b=", b, "d=", d) //b=[1 200 300] d=[1 200 300]
}
