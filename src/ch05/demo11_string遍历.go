package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱Go语言"
	fmt.Println("字符串长度：", len(s))
	fmt.Println("----------------------------")

	//for ... range遍历字符串
	for i, ch := range s {
		fmt.Printf("%d:%c \n", i, ch)
	}
	fmt.Println("----------------------------")

	//遍历所有字节
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x \n", i, ch)
	}
	fmt.Println("----------------------------")

	//遍历所有字符
	count := 0
	for i, ch := range []rune(s) {
		fmt.Printf("%d:%c \n", i, ch)
		count++
	}
	fmt.Println("字符串长度：", count)
	fmt.Println("字符串长度：", utf8.RuneCountInString(s))
	fmt.Println("----------------------------")

}
