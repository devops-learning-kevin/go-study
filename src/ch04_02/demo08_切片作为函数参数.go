package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T,%v\n", &arr, &arr) //*[]int,&[1 2 3 4 5]

	fmt.Println("调用函数前arr的值", arr) //调用函数前arr的值 [1 2 3 4 5]
	modify(&arr)
	fmt.Println("调用函数后arr的值", arr) //调用函数后arr的值 [250 2 3 4 5]
}

func modify(arr *[]int) {
	(*arr)[0] = 250
}
