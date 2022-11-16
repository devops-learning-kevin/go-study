package main

/*
基本数据类型（原生数据类型）：整形、浮点型、布尔型、字符串、字符（byte,rune）
复合数据类型（派生数据类型）：指针（pointer）、数组（array）、切片（slice）、映射（map）、
函数（function)、结构体（struct）、通道（channel）
*/
import "fmt"

func main() {
	a := 100
	var b byte = 100
	var c rune = 200
	var e byte = 'a' //byte = uint8
	var f rune = '一' //unicode 字符集
	var str string = "abcdef"

	fmt.Printf("%T,%v\n", a, a)
	fmt.Printf("%T,%v\n", b, b)
	fmt.Printf("%T,%v\n", e, e)
	fmt.Printf("%T,%v\n", f, f)
	fmt.Printf("%T,%v\n", c, c)
	fmt.Printf("%T,%v\n", str, str)

	var temp = `
	x := 10
	y := 20
	z := 30
	fmt.Println(x, "  ", y, "  ", z)
	x, y, z = y, z, x
	fmt.Println(x, "  ", y, "  ", z)
    `
	fmt.Println(temp)

}
