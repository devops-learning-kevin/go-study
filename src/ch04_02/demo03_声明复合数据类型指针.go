package main

import "fmt"

type Student struct {
	name    string
	age     int
	married bool
	sex     int8
}

func main() {
	s1 := Student{"Kevin", 43, true, 1}
	s2 := Student{"Lucy", 36, false, 0}

	//var a *Student=&s1
	//var b *Student &s2

	a := &s1
	b := &s2

	fmt.Println("------------------------------------")
	fmt.Printf("s1的类型为：%T,值为%v \n", s1, s1)
	fmt.Printf("s2的类型为：%T,值为%v \n", s2, s2)

	fmt.Println("------------------------------------")
	fmt.Printf("a的类型为：%T,值为%v \n", a, a)
	fmt.Printf("b的类型为：%T,值为%v \n", b, b)

	fmt.Println("------------------------------------")
	fmt.Printf("*a的类型为：%T,值为%v \n", *a, *a)
	fmt.Printf("*b的类型为：%T,值为%v \n", *b, *b)

	//对应结构体对象的成员，可以用结构体对象变量名.成员名访问，也可以使用结构体的指针.成员名访问
	fmt.Println("------------------------------------")
	fmt.Println(s1.name, s1.age, s1.married, s1.sex)
	fmt.Println(a.name, a.age, a.married, a.sex)

	fmt.Println("------------------------------------")
	fmt.Println((*a).name, (*a).age, (*a).married, (*a).sex)
	fmt.Println(&a.name, &a.age, &a.married, &a.sex) //成员变量的地址
	fmt.Println(*&a.name, *&a.age, *&a.married, *&a.sex)

}
