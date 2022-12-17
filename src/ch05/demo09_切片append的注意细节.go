package main

import (
	"fmt"
	"strconv"
)

func main() {
	//sa := []string{}  //这样写也不是不可以,不建议
	//var sa []string  //这样定义的空切片，随着append新元素，地址会重新分配
	//sa := make([]string, 0, 0)  //虽然使用了make分配了初始内存地址，sa初始不为nil了，但没分配足够的空间，所以扩充的过程中，内存还是要重新分配
	sa := make([]string, 0, 20) //分配了足够的内存空间，内存地址也确定了，所以以后扩充过程中，不会再重新分配内存，效率必定是最优的

	if sa == nil {
		fmt.Println("s==nil", "len = ", len(sa), "cap = ", cap(sa))
	}

	for i := 0; i < 12; i++ {
		sa = append(sa, strconv.Itoa(i))
		printSliceData(sa)
	}
}

func printSliceData(s []string) {
	fmt.Printf("地址：%p,len:%d,cap:%d,value:%v\n", s, len(s), cap(s), s)
}
