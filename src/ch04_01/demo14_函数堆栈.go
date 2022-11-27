package main

import "fmt"

func main() {
	retmain := fn1(1)
	fmt.Println(retmain)
}

func fn1(a1 int) (ret1 int) {
	a12 := a1 + 1
	ret1 = fn2(a12)
	return
}

func fn2(a2 int) (ret2 int) {
	a23 := a2 + 2
	ret2 = fn3(a23)
	return
}

func fn3(a3 int) (ret3 int) {
	a34 := a3 + 3
	ret3 = fn4(a34)
	return
}

func fn4(a4 int) (ret4 int) {
	ret4 = a4 + 4
	return
}
