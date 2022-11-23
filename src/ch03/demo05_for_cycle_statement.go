package main

import "fmt"

func main() {
	//oneAddToHundred(100)
	baseFor()
}
func oneAddToHundred(maxnum int) {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum = sum + i
		fmt.Printf("%d + %d = %d\n", sum-i, i, sum)
	}
}

func baseFor() {
	//for 循环基本语法结构第一种形式
	for i := 0; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n------------------------ \n")

	// 第二种形式
	i := 0
	for ; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n------------------------ \n")

	//第三种形式
	for ; ; i++ {
		if i > 20 {
			break
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n------------------------ \n")

	//第四种形式
	for {
		if i > 40 {
			break
		}
		i++
		fmt.Printf("%d ", i)

	}

}
