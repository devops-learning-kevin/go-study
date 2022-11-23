package main

import "fmt"

/*
1、定义一个函数类型
2、实现定义的函数类型
3、作为参数调用
*/

type myFunc func(int) bool
type myFunc2 func(int) int

//判断元素是否偶数
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

//判断元素是否奇数
func isOdd(num int) bool {
	if num%2 == 0 {
		return false
	} else {
		return true
	}
}

//求整型切片的所有元素的平方
func pingfang(num int) int {
	return num * num
}

//传入切片和方法名来过滤或者处理切片
func FilterArr(arr []int, f myFunc) []int {
	var result []int
	for _, value := range arr {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

//传入切片和方法名来过滤或者处理切片
func FilterArr2(arr []int, f myFunc2) []int {
	var result []int
	for _, value := range arr {
		result = append(result, f(value))
	}
	return result
}

func main() {
	arr := []int{12, 375, 98, 32, 4, 7, 289, 403, 18, 20}
	fmt.Println("slice = ", arr)

	//过滤出奇数
	odd := FilterArr(arr, isOdd)
	fmt.Println("奇数：", odd)

	//过滤出偶数
	even := FilterArr(arr, isEven)
	fmt.Println("偶数：", even)

	//平方
	pingf := FilterArr2(arr, pingfang)
	fmt.Println("平方：", pingf)

}
