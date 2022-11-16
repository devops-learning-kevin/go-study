package main

import "fmt"

var a = 21.0
var b = 5.0
var c float64

func main() {
	//Arithmetic()
	//Relational()
	//logical()
	//Bitwise()
	Assignment()
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

	a = 5
	b = 20
	if a <= b {
		fmt.Printf("第四行 - a 小于等于 b \n")
	} else {
		fmt.Printf("第四行 - a 大于 b \n")
	}

	if a >= b {
		fmt.Printf("第五行 - a 大于等于 b \n")
	} else {
		fmt.Printf("第五行 - a 小于 b \n")
	}

}

//逻辑运算符
func logical() {
	a := true
	b := false

	if a && b {
		fmt.Printf("第一行 - 条件为 true\n")
	}

	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a和b的值*/
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n") //fasle
	}

	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n") //true
	}

}

//位运算符
func Bitwise() {
	fmt.Println(252 & 63)
	fmt.Println(178 | 94)
	fmt.Println(20 ^ 5)
	fmt.Println(3 << 4)
	fmt.Println(11 >> 2)
	fmt.Println(11 >> 4)
}

//赋值运算符
func Assignment() {
	c = a
	fmt.Printf("第1行==运算符实例，c值为=%f\n", c)
	c += a
	fmt.Printf("第2行=+=运算符实例，c值为=%f\n", c)
	c -= a
	fmt.Printf("第3行--=运算符实例，c值为=%f\n", c)
	c *= a
	fmt.Printf("第4行-*=运算符实例，c值为=%f\n", c)
	c /= a
	fmt.Printf("第5行-/=运算符实例，c值为=%f\n", c)

	d := 200
	d <<= 2
	fmt.Printf("第6行- <<=运算符实例，d值为=%d\n", d) //d = d*2^2 = 800

	d >>= 2
	fmt.Printf("第7行- >>=运算符实例，d值为=%d\n", d) //d = d/2^2 = 200

	d &= 2
	fmt.Printf("第8行- &=运算符实例，d值为=%d\n", d) // d = d & 2 = 0

	d |= 2
	fmt.Printf("第9行- |=运算符实例，d值为=%d\n", d) // d = d | 2 = 2

	d ^= 2
	fmt.Printf("第10行- ^=运算符实例，d值为=%d\n", d) // d = d^2 = 0
}
