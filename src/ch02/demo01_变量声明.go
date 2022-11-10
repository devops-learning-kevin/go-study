package main

import "fmt"

var m = "19"

func main() {

	var (
		a int
		b string
		c float32
		d []int
		e int32 = 100
		f [3]string
		g bool
		h func() string
		l = 200.234
	)

	m := "19"
	n := "190"
	m, n, q := "a", "b", 10  //q 为新变量，可以初始化定义模式赋值
	m, _, r := "c", "d", 100 //匿名变量
	fmt.Printf("%T,%v \n", a, a)
	fmt.Printf("%T,%q,%v \n", b, b, b)
	fmt.Printf("%T,%v \n", c, c)
	fmt.Printf("%T,%v \n", d, d)
	fmt.Printf("%T,%v \n", e, e)
	fmt.Printf("%T,%v \n", f, f)
	fmt.Printf("%T,%v \n", g, g)
	fmt.Printf("%T,%v \n", h, h)
	fmt.Printf("%T,%v \n", l, l)
	fmt.Printf("%T,%v \n", m, m)
	fmt.Printf("%T,%v \n", n, n)
	fmt.Printf("%T,%v \n", q, q)
	fmt.Printf("%T,%v \n", r, r)
	fmt.Println(a, b, c, d, e, f, g, h, l)

	// 变量直接交换，程序员期待已久的语法
	x := 10
	y := 20
	z := 30
	fmt.Println(x, "  ", y, "  ", z)
	x, y, z = y, z, x
	fmt.Println(x, "  ", y, "  ", z)
}
