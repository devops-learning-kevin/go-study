package main

import "fmt"

/* 打印格式化
(一)、通用
   %v   值的默认格式表示  value
   %+v  类似%v，但输出结构体时，会添加字段名
   %#v  值的Go语法表示
   %T   值的类型的Go语法表示  type

(二)、布尔值
   %t   单词的true或false  true

(三)、整数
   %b   表示为二进制  binary
   %c   该值对应的unicode码值  char
   %d   表示为十进制   digital
   %8d  表示该整型长度是8,不足8则数值钱补空格。如果超出8，则以实际为准。
   %08d 表示该整型长度是8，不足8位的，在数字前面补0。如果超出8，则以实际为准。
   %o   表示为8进制  octal
   %q   该值对应的引号括起来的go语法字符字面值，必要时会采用安全的转义表示  quotation
   %x   表示为十六进制，使用a-f  hex
   %X   表示为十六进制，使用A-F
   %U   表示为Unicode格式： U+1234,等价于 "U+%04X"  unicode

(四)、浮点数与复数的两个组分
   %b   无小数部分、二进制指数的科学计数法，如  -123456p-78; 参见strconv.FormatFloat
   %e   (=%.6e) 有6位小数部分的科学计数法，如 -1234.456e+78
   %E   科学计数法，如 -1234.456E+78
   %f   (=%.6f) 有6位小数部分，如123.456123   float
   %F   等价于 %F
   %g   根据实际情况采用%e或%f格式 （以获得更简洁、准确的输出）
   %G   根据实际情况采用%E或%F格式 （以获得更简洁、准确的输出）

(五)、字符串和byte
   %s   直接输出字符串或者[]byte   string
   %q   该值对应的引号括起来的go语法字符字面值，必要时会采用安全的转义表示  quotation
   %x   每个字节用两个字符十六进制数表示（使用a-f）
   %F   每个字节用两个字符十六进制数表示（使用A-F）

(六)、指针
   %p  表示为十六进制，并加上前导的ox pointer
   没有 %u。整数如果是无符号类型自然输出也是无符号的。类似的，也没有必要指定操作数的尺寸（int8,int64)。
   宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不作填充。精度通过(可选的)宽度后跟点号后跟的十进制数指定。如果未指定精度，会使用默认精度;如果点号后没有跟数字，表示精度为0。举例如下:
    %f:默认宽度，默认精度
    %9f宽度9，默认精度
    %.2f默认宽度，精度2
    %9.2f宽度9，精度2
    %9.f宽度9，精度0

*/

type Student struct {
	x, y int
}

func main() {
	// 通用格式
	str := "steven"
	fmt.Printf("%T,%v\n", str, str)

	var a rune = '一'
	fmt.Printf("%T,%v \n", a, a)

	p := Student{1, 2}
	fmt.Printf("%T,%v\n", p, p)

	// 布尔格式
	fmt.Printf("%T,%t \n", true, true)
	fmt.Printf("%T,%v \n", true, true)

	// 整型格式
	fmt.Printf("%T,%d \n", 123, 123)
	fmt.Printf("%T,%5d \n", 123, 123)
	fmt.Printf("%T,%05d \n", 123, 123)
	fmt.Printf("%T,%05d \n", 123, 123456)
	fmt.Printf("%T,%b \n", 123, 123)

	// 进制转换
	strb := fmt.Sprintf("%b", 123)
	fmt.Println(strb)

	fmt.Printf("%x \n", 123)
	fmt.Printf("%X \n", 123)
	fmt.Printf("%U \n", 'a')
	fmt.Printf("%U \n", '一')

	//浮点型
	fmt.Printf("%f \n", 123.1) // %f = %.2f
	fmt.Printf("%.2f \n", 123.1)
	fmt.Printf("%.2f \n", 123.12356) //四舍五入切割
	fmt.Printf("%.2f \n", 123.12556)
	fmt.Printf("%e \n", 123.12556) //科学计数
	fmt.Printf("%.10e \n", 123.12556)
	fmt.Printf("%.1e \n", 123.12556)

	//字符串
	fmt.Printf("%s \n", "学习区块链")
	fmt.Printf("%q \n", "学习区块链")
	fmt.Printf("%s \n", "")
	fmt.Printf("%q \n", "")

	//字节数组
	arr := [3]byte{97, 98, 99}
	fmt.Printf("%T,%s \n", arr, arr)

	arr1 := []rune{'a', 'b', 'c', '一'}
	fmt.Printf("%T,%s \n", arr1, arr1)

	arr2 := []byte{'a', 'b', 'c', 'A'}
	fmt.Printf("%T,%s \n", arr2, arr2)

	arr3 := []byte{'a', 'b', 'c', 'A'}
	fmt.Printf("%T,%x \n", arr3, arr3)

}
