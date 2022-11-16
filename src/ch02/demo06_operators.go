package main

import "fmt"

var a = 21.0
var b = 5.0
var c float64

func main() {
	Arithmetic()
	Relational()
}

//算术运算符
func Arithmetic() {
	c = a + b
	fmt.Printf("第一行 - c 的值为 %.2f\n", c)

	c = a - b
	fmt.Printf("第二行 - c 的值为 %.2f\n", c)

	c = a * b
	fmt.Printf("第三行 - c 的值为 %.2f\n", c)

	c = a / b
	fmt.Printf("第四行 - c 的值为 %.2f\n", c)

	//c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", int(a)%int(b))

	a++
	fmt.Printf("第六行 - a 的值为 %f\n", a)

	a = 21
	a--
	fmt.Printf("第七行 - a 的值为 %f\n", a)
}

//关系运算符
func Relational() {
	if a == b {
		fmt.Printf("第一行 - a 等于 b \n")
	} else {
		fmt.Printf("第一行 - a 不等于 b \n")
	}

	if a < b {
		fmt.Printf("第二行 - a 小于 b \n")
	} else {
		fmt.Printf("第二行 - a 不小于 b \n")
	}

	if a > b {
		fmt.Printf("第三行 - a 大于 b \n")
	} else {
		fmt.Printf("第三行 - a 不大于 b \n")
	}

}
