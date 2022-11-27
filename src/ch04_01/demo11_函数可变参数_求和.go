package main

import "fmt"

func main() {
	//1、传多个值
	fmt.Println(AddSum(1, 2, 3, 4, 5, 6, 7))

	//2、传0个值
	fmt.Println(AddSum())

	//3、传切片作为参数
	sum := []int{3, 5, 7, 10, 15}
	fmt.Println(AddSum(sum...))

}

//累加求和，参数不定，参数个数从0-n
func AddSum(nums ...int) (sum int) {
	fmt.Printf("可变参数类型打印：%T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return
}
