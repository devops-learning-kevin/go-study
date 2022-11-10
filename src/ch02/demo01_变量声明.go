package main

import "fmt"

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
	fmt.Printf("%T,%v \n", a, a)
	fmt.Printf("%T,%q,%v \n", b, b, b)
	fmt.Printf("%T,%v \n", c, c)
	fmt.Printf("%T,%v \n", d, d)
	fmt.Printf("%T,%v \n", e, e)
	fmt.Printf("%T,%v \n", f, f)
	fmt.Printf("%T,%v \n", g, g)
	fmt.Printf("%T,%v \n", h, h)
	fmt.Printf("%T,%v \n", l, l)
	fmt.Println(a, b, c, d, e, f, g, h, l)
}
