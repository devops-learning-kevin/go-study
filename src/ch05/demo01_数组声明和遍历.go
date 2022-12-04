package main

import "fmt"

//声明数组的方式1
var arr [3]int
var arr2 = [4]int{1, 2, 3, 4}
var arr3 [5]int
var arr4 []int

func main() {
	//声明数组的方式2
	a := [4]float64{2.3, 3, 3.14, 18.8}
	fmt.Println(a)

	//未指定长度的数组，使用[...]
	b := [...]int{2, 3, 4}

	//遍历数组的方式1
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}

	fmt.Println()

	//遍历数组的方式2
	for _, v := range b {
		fmt.Print(v, "\t")

	}

	fmt.Println()

	//数组即使没有被初始化，也不是nil
	//if arr3 == nil{
	//	fmt.Print("arr3 == nil")
	//}

	//切片如果没有被初始化，是nil
	if arr4 == nil {
		fmt.Print("arr4 == nil\n")
	}

	//遍历arr空数组
	fmt.Println("遍历arr[3]空数组:", arr)
	for _, v := range arr {
		fmt.Print(v, "\t")
	}

}
