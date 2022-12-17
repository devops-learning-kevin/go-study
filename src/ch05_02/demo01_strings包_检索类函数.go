package main

import (
	"fmt"
	"strings"
	"unicode"
)

/*

(一)、检索字符串
1、func Contains(s, substr string) bool 判断字符串s是否包含substr字符串

2、func ContainsAny(s, chars string) bool 判断字符串s是否包含chars字符串中的任一字符

3、func ContainsRune(s string, r rune) bool 判断字符串s是否包含unicode码值r

4、func Count(s, sep string) int 返回字符串s包含字符串sep的个数

5、func HasPrefix(s, prefix string) bool 判断字符串s是否有前缀字符串prefix

6、func HasSuffix(s, suffix string) bool判断字符串s是否有后缀字符串suffix

7、func Index(s, sep string) int 返回字符串s中字符串sep首次出现的位置

8、func IndexAny(s, chars string)int 返回字符串chars中的任一unicode码值r在s中首次出现的位置

9、func IndexByte(s string, c byte) int 返回字符串s中字符c首次出现位置

10、func IndexFunc(s string,f func(rune) bool)int 返回字符串s中满足函数f(r)==true字符首次出现的位置

11、func IndexRune(s string, r rune)int 返回unicode码值r在字符串中首次出现的位置

12、func LastIndex(s,sep string)int 返回字符串s中字符串sep最后一次出现的位置

13、func LastIndexAny(s,chars string) int 返回字符串s中任意一个unicode码值r最后一次出现的位置

14、func LastIndexByte(s string,c byte) int 返回字符串s中字符c最后一次出现的位置

15、func LastIndexFunc(s string,f func(rune) bool)int 返回字符串s中满足函数f(r)==true字符最后一次出现的位置
*/

func main() {
	//TestContains()
	//TestContainsAny()
	//TestContainsRune()
	//TestCount()
	//TestHasPrefix()
	//TestHasSuffix()
	//TestIndex()
	//TestIndexAny()
	//TestIndexByte()
	//TestIndexRune()
	//TestIndexFunc()
	//TestLastIndex()
	//TestLastIndexAny()
	//TestLastIndexByte()
	TestLastIndexFunc()
}

func TestContains() {
	fmt.Println(strings.Contains("seafood", "foo")) //true
	fmt.Println(strings.Contains("seafood", "bar")) //false
	fmt.Println(strings.Contains("seafood", ""))    //true
	fmt.Println(strings.Contains("", ""))           //true
	fmt.Println(strings.Contains("keivn锟", "锟"))    //true
}

func TestContainsAny() {
	fmt.Println(strings.ContainsAny("team", "i"))        //false
	fmt.Println(strings.ContainsAny("failure", "u & i")) //true
	fmt.Println(strings.ContainsAny("foo", ""))          //false
	fmt.Println(strings.ContainsAny("", ""))             //fasle

}

func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁丂", '丁'))   //true
	fmt.Println(strings.ContainsRune("一丁丂", 19969)) //true

}

func TestCount() {
	fmt.Println(strings.Count("cheese", "e")) //3
	fmt.Println(strings.Count("one", ""))     //4

}

func TestHasPrefix() {
	fmt.Println(strings.HasPrefix("1000phone news", "1000"))  //true
	fmt.Println(strings.HasPrefix("1000phone news", "1000a")) //false
}

func TestHasSuffix() {
	fmt.Println(strings.HasSuffix("1000phone news", "news")) //true
	fmt.Println(strings.HasSuffix("1000phone news", "new"))  //false
}

func TestIndex() {
	fmt.Println(strings.Index("chicken", "ken")) //4
	fmt.Println(strings.Index("chicken", "dmr")) //-1
}

func TestIndexAny() {
	fmt.Println(strings.IndexAny("abcABC120", "教育基地A")) //3

}

func TestIndexByte() {
	fmt.Println(strings.IndexByte("123abc", 'a')) //3
}

func TestIndexRune() {
	fmt.Println(strings.IndexRune("abcABC120", 'C')) //5
	fmt.Println(strings.IndexRune("It培训教育", '教'))    //8

}

func TestIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello,世界！", f)) //6 第一个汉字出现的索引位置

}

func TestLastIndex() {
	fmt.Println(strings.LastIndex("abcdefgabcdefg", "de")) //10
	fmt.Println(strings.Index("go gopher", "go"))          //0
	fmt.Println(strings.LastIndex("go gopher", "go"))      //3
	fmt.Println(strings.LastIndex("go gopher", "rodent"))  //-1

}

func TestLastIndexAny() {
	fmt.Println(strings.LastIndexAny("chicken", "aeiouy")) // 5
	fmt.Println(strings.LastIndexAny("crwth", "aeiouy"))   //-1

}

func TestLastIndexByte() {
	fmt.Println(strings.LastIndexByte("abcABCA123", 'A')) //6

}

func TestLastIndexFunc() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.LastIndexFunc("Hello,世界！", f))      //9 第一个汉字出现的索引位置
	fmt.Println(strings.LastIndexFunc("Hello wold,中国！", f)) //14 第一个汉字出现的索引位置

}
