package main

import (
	"fmt"
	"strings"
)

/*
(五)、比较字符串
1、func Compare(a,b string)int
按字典顺序比较a和b字符串大小

2、func EqualFold(s,t string) bool
判断s和t两个utf8字符串是否相等，忽略大小写

3、func Repeat(s string, count int) string
将字符串s重复count次返回

4、func Replace(s, old, new string, n int) string
替换字符串s中old字符为new字符并返回，n<0是替换所有old字符串

5、func Join(a [string, sep string) string
将a中的所有字符串连接成一个字符串，使用字符串sep作为分隔符
*/

func main() {
	TestCompare()
	TestEqualFold()
	TestRepeat()
	TestReplace()
	TestJoin()
}
func TestCompare() {
	fmt.Println(strings.Compare("abc", "bcd")) //-1
	fmt.Println("abc" < "bcd")
}

func TestEqualFold() {
	fmt.Println(strings.EqualFold("Go", "go")) //true
}

func TestRepeat() {
	fmt.Println("g" + strings.Repeat("o", 8) + "le") //goooooooole
}

func TestReplace() {
	fmt.Println(strings.Replace("王老大 王老二 王老三", "王", "张", 2)) //张老大 张老二 王老三
	fmt.Println(strings.Replace("王老大王老二王老三", "王", "张", -1))  //张老大 张老二 张老三
}

func TestJoin() {
	s := []string{"abc", "ABC", "123"}
	fmt.Println(strings.Join(s, ",")) //abc,ABC,123
	fmt.Println(strings.Join(s, ""))  //abcABC123
}
