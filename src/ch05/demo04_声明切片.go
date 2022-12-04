package main

import "fmt"

//声明切片的方式1
var s1 []int

//声明切片的方式2.1
var s2 = make([]int, 5)    //指定切片的长度为5
var s3 = make([]int, 5, 7) //指定切片的长度为5，最大容量为7

func main() {
	//声明切片的方式2.2 直接:=
	s4 := make([]int, 5)
	s5 := make([]int, 5, 7)
	s6 := make([]int, 0)
	printSliceMsg(s1)
	printSliceMsg(s2)
	printSliceMsg(s3)
	printSliceMsg(s4)
	printSliceMsg(s5)

	if s1 == nil {
		fmt.Println("s1==nil", "len(s1)=", len(s1)) //切片是nil，但是可以取长度
	}

	if s2 == nil {
		fmt.Println("s2==nil", "len(s2)=", len(s2)) //make分配了内存，指定了长度的切片不是nil
	} else {
		fmt.Println("s2!=nil", "len(s2)=", len(s2))
	}

	//make长度为0的切片不为nil
	if s6 == nil {
		fmt.Println("s6==nil", "len(s6)=", len(s6))
	} else {
		fmt.Println("s6!=nil", "len(s6)=", len(s6))
	}

	//fmt.Printf("len=%d,cap=%d,slice=%v \n", len(s1), cap(s1), s1)
	//fmt.Printf("len=%d,cap=%d,slice=%v \n", len(s2), cap(s2), s2)
	//fmt.Printf("len=%d,cap=%d,slice=%v \n", len(s3), cap(s3), s3)
	//fmt.Printf("len=%d,cap=%d,slice=%v \n", len(s4), cap(s4), s4)
	//fmt.Printf("len=%d,cap=%d,slice=%v \n", len(s5), cap(s5), s5)

}

//把打印封装成函数，可以提高复用性，代码更简洁
func printSliceMsg(s []int) {
	fmt.Printf("len=%d,cap=%d,slice=%v \n", len(s), cap(s), s)
}
