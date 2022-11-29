package main

import "fmt"

func main() {
	//实际变量
	a := 20

	//声明指针变量
	var ip *int

	//给指针变量赋值，将变量a的地址赋值给ip
	ip = &a

	//打印a的类型和值
	fmt.Printf("a的类型是%T,值是%v\n", a, a) //int 20

	//打印&a的类型和值
	fmt.Printf("&a的类型是%T,值是%v\n", &a, &a)

	//打印ip的类型和值
	fmt.Printf("ip的类型是%T,值是%v\n", ip, ip)

	//打印*ip的类型和值
	fmt.Printf("*ip的类型是%T,值是%v\n", *ip, *ip) //int 20

	//打印*&a的类型和值
	fmt.Printf("*&a的类型是%T,值是%v\n", *&a, *&a) // int 20

	fmt.Println(a, &a, *&a)
	fmt.Println(ip, &ip, *ip, *(&ip), &(*ip))

}
