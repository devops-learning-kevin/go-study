package main

import "fmt"

func main() {
	//1、声明map的方式1
	var map1 map[string]string //没有分配内存空间，所以是nil
	printMapInfo(map1)

	//2、声明map的方式2
	map2 := make(map[string]string) //分配了内存空间，所以不是nil
	printMapInfo(map2)

	//3、map的key可以是：int、float、bool、string、数组
	//一定不可以是:切片、函数、map

	var m1 map[int]string
	var m2 map[float64]string
	var m3 map[string]string
	//var m4 map[[]int]string //nvalid map key type: comparison operators == and != must be fully defined for the key type
	var m4 map[bool]string
	var m5 map[[6]int]string
	//var m6 map[map[string]string]string //Invalid map key type: comparison operators == and != must be fully defined for the key type

	fmt.Println(m1, m2, m3, m4, m5)

	fmt.Println(len(m1))
	//fmt.Println(cap(m1))  //Invalid argument for the cap function

}

func printMapInfo(m map[string]string) {
	fmt.Println("-----------------------------")
	fmt.Printf("%T,%v,len(m)=%d,m==nil?%v \n", m, m, len(m), m == nil)
}
