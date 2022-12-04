package main

import "fmt"

func main() {
	//创建并初始化切片
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSliceInfo(numbers)

	numbers1 := numbers[:] //[0 1 2 3 4 5 6 7 8]
	printSliceInfo(numbers1)

	numbers2 := numbers[1:4] //[1 2 3]
	printSliceInfo(numbers2)

	numbers3 := numbers[:3] //[0 1 2]
	printSliceInfo(numbers3)

	numbers4 := numbers[4:] //[4 5 6 7 8]
	printSliceInfo(numbers4)

}

func printSliceInfo(x []int) {
	fmt.Printf("len=%d,cap=%d,slice=%v\n", len(x), cap(x), x)
}
