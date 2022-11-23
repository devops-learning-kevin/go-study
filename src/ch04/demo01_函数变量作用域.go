package main

import "fmt"

//声明全局变量
var a int = 7
var b int = 9

func main() {
	//声明局部变量
	a, b, c := 10, 20, 0
	fmt.Printf("main()函数中a= %d\n", a) //10
	fmt.Printf("main()函数中b= %d\n", b) //20
	fmt.Printf("main()函数中c= %d\n", c) //0
	c, _ = sum(a, b)
	fmt.Printf("main()函数中c= %d\n", c) //33

	d, _ := sum2(a, b)
	fmt.Printf("main()函数中d= %d\n", d)

}

//两个数值相加
func sum(a, b int) (int, int) {
	a++        //11
	b += 2     //22
	c := a + b //33
	d := a * b
	fmt.Printf("sum()函数中a=%d\n", a) //11
	fmt.Printf("sum()函数中b=%d\n", b) //22
	fmt.Printf("sum()函数中c=%d\n", c) //33
	fmt.Printf("sum()函数中d=%d\n", d)
	return c, d
}

func sum2(a, b int) (c, d int) {
	a++       //11
	b += 2    //22
	c = a + b //33
	d = a * b
	fmt.Printf("sum2()函数中a=%d\n", a) //11
	fmt.Printf("sum2()函数中b=%d\n", b) //22
	fmt.Printf("sum2()函数中c=%d\n", c) //33
	return
}
