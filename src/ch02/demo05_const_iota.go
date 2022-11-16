package main

import "fmt"

//常量
type Man struct {
	age, height int
}

func main() {
	const NAME string = "kevin"
	//NAME = "sunwukong"	  //Cannot assign to NAME

	//结构体类型不能作为常量
	//man := Man{30,177}
	//const PERSON = man   //Const initializer 'man' is not a constant

	//常量组,第一个常量必须赋值
	//常量组没赋值的成员，值与上一个常量保持一致
	const (
		//X   //Missing value in the const declaration
		X = "A"
		Y
		A = 10
		B
		C
	)
	fmt.Println(X, Y, A, B, C) //A A 10 10 10

	//常量组计数器 iota
	const (
		L = iota //0
		M        //1
		N        //2
	)
	fmt.Println(L, M, N)

	const (
		L1 = iota //0
		M1
		N1
	)
	fmt.Println(L1, M1, N1)

	const (
		L2 = "KEVIN" // iota = 0
		M2 = iota    // iota = 1
		N2           //iota = 2
	)
	fmt.Println(L2, M2, N2)

	// 注意常量组中，未赋值的常量的赋值应该是上一个赋值的常量的表达式，而不是计算的最终值
	const (
		i = 1 << iota //1  1*2^iota
		j = 3 << iota //6
		k             //12   k = 3 << iota
		l             //24   l = 3 << iota
	)
	fmt.Println(i, j, k, l)

	const (
		i1 = 1 << iota //1  1*2^iota
		j1
		k1
		l1
	)
	fmt.Println(i1, j1, k1, l1)

	const (
		a1 = '一'
		b1
		c1 = iota
		d1
	)
	fmt.Println(a1, b1, c1, d1) //19968，19968，2，3

	const name = iota + 1
	fmt.Printf("%T,%v\n", name, name)

}
