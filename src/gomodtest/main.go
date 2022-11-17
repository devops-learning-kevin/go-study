package main

import (
	"fmt"
	"github.com/jinzhu/configor"
	"sum"
)

func main() {
	fmt.Println("hello world")
	fmt.Println("sum=", sum.Sum(2, 5))
	fmt.Println("使用外部包测试：", configor.Config{})
}
