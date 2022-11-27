package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("i=%d \t", i)
		fmt.Println(add(i))
	}
}

func add(num int) int {
	sum := 0
	sum += num
	return sum
}
