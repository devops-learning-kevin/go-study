package main

import "fmt"

func main() {
	summation()
	summation2()
	cutBamboo()
	traverseString()
	traverseSlice()
}

//计算1-100的和
func summation() {
	result := 0
	for i := 0; i <= 100; i++ {
		result += i
	}
	fmt.Println(result)
}

//计算1-100之间3的倍数的和
func summation2() {
	result := 0
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			result += i
			fmt.Print(i)
			if i < 99 {
				fmt.Print("+")
			} else {
				fmt.Printf("=%d\n", result)
			}
		}
		i++
	}
	//fmt.Println(result)

}

// 3、截竹竿。32米的竹竿，每次截取1.5米，最快截取几次之后能小于4米
func cutBamboo() {
	count := 0
	for i := 32.0; i > 4; i -= 1.5 {
		count++
	}
	fmt.Println(count)

}

//4、遍历字符串，获得字符
func traverseString() {
	str := "123ABCabc一丁丂"
	for i, value := range str {
		fmt.Printf("第%d位的字符值是: %v,原字符是：%c\n", i, value, value)
	}

	//utf-8、 utf-16
}

// 5、遍历一个切片

func traverseSlice() {
	arr := []int{100, 200, 300}
	for i, value := range arr {
		fmt.Println(i, ":", value)
	}

}
