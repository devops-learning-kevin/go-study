package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*
1、func Fields(s string) []string
将字符串s以空白字符分割，返回一个切片

2、func FieldsFunc(s string,f func(rune) bool) []string
将字符串s以满足f(r)==true的字符分割，返回一个切片

3、func Split(s, sep string) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后去掉sep

4、func SplitAfter(s,sep string) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep

5、func SplitAfterN(s, sep string，n int) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后附上sep，n决定返回的切片数

6、func SplitN(s, sep string, n int) []string
将字符串s以sep作为分隔符进行分割，分割后字符最后去掉sep，n决定返回的切片数
*/

func main() {
	TestFields()
	TestFieldsFunc()
	TestSplit()
	TestSplitN()
	TestSplitAfter()
	TestSplitAfterN()

}

func TestFields() {
	fmt.Println(strings.Fields("abc 123 ABC xyz XYZ")) //[abc 123 ABC xyz XYZ]
}

func TestFieldsFunc() {
	f := func(c rune) bool {
		//return c == '='
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	//fmt.Println(strings.FieldsFunc("abc=123ABC=xyz=XYZ", f)) //[abc 123ABC xyz XYZ]
	fmt.Println(strings.FieldsFunc("abc&123%ABC#xyz*XYZ", f)) //[abc 123 ABC xyz XYZ]

}

func TestSplit() {
	fmt.Printf("%v\n", strings.Split("a,b,c", ","))                        //[a b c]
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))                        //["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a ")) //["" "man " "plan " "canal panama"]
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))                         //["" "x" "y" "z" ""]
	fmt.Printf("%q\n", strings.Split("", "xyz"))                           //[""]

}

func TestSplitN() {
	fmt.Printf("%q\n", strings.SplitN("x,y,z", ",", 2)) //["x" "y,z"]
	fmt.Printf("%q\n", strings.SplitN("x,y,z", ",", 1)) //["x,y,z"]

}

func TestSplitAfter() {
	fmt.Printf("%v\n", strings.SplitAfter("a,b,c", ",")) //[a b c]

}

func TestSplitAfterN() {
	fmt.Printf("%q\n", strings.SplitAfterN("x,y,z", ",", 2)) //["x," "y,z"]

}
