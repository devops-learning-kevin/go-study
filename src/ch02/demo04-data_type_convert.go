package main

import "fmt"

// 数据类型转换
func main() {
	chinese := 90
	english := 80.9
	//整型与浮点型的相互转换, 浮点型默认与操作系统的位数有关，现在大部分操作系统是64位，所以给一个float值，默认是float64
	avg := (chinese + int(english)) / 2
	//avg2 := (float32(chinese) + english) / 2   //Invalid operation: float32(chinese) + english (mismatched types float32 and float64)
	avg2 := (float64(chinese) + english) / 2
	fmt.Printf("%T,%v\n", avg, avg)
	fmt.Printf("%T,%q\n", avg, avg)
	fmt.Printf("%T,%v\n", avg2, avg2)
	fmt.Printf("%T,%q\n", avg2, avg2)

	//bool型不能与其它类型转换
	//flag := true
	//string(flag)  //Cannot convert an expression of the type 'bool' to the type 'string'

	//字符串不能转换到整型
	//str := "12356"
	//int(str)    //Cannot convert an expression of the type 'string' to the type 'int'

	//字符串不能转换为整型，但整型是可以转换为字符串类型,
	//但这个整型转换字符串，并不是直接原样字面转换，而是按照unicode编码转换

	res := string(chinese)
	fmt.Printf("%T,%v\n", res, res) //为什么是大Z，不是"90"

	res1 := string(123456789)
	fmt.Printf("%T,%q\n", res1, res1)

	//float不能转换为字符串
	//res2:= string(english)   //Cannot convert an expression of the type 'float64' to the type 'string'
	//fmt.Printf("%T,%v\n", res2, res2)

	res3 := string(19968)
	fmt.Printf("%T,%q\n", res3, res3)

	x := 'Z' //字符类型和整型通用
	fmt.Printf("%T,%q\n", x, x)

	// go语言中没有 char类型 char类型都用整型替代。
	//var cc char = 'm'

}
