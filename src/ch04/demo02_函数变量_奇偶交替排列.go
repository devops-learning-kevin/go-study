package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ABCdEFghigkLmnopQRsT"
	//result := processLetter(str)
	result := SringToCase(str, processLetter)
	fmt.Println(result)

	fmt.Println(SringToCase2(str, processLetter))
}

// 处理字符串，实现大小写字母奇偶交替出现
func processLetter(str string) string {
	result := ""
	for i, value := range str {
		if i%2 == 0 {
			result += strings.ToUpper(string(value))
		} else {
			result += strings.ToLower(string(value))
		}
	}
	return result
}

// 函数值作为参数
func SringToCase(str string, myfunc func(string) string) string {
	return myfunc(str)
}

// 定义函数类型
type caseFunc func(string) string

//函数类型作为函数的参数
func SringToCase2(str string, myfunc caseFunc) string {
	return myfunc(str)
}
