package main

import "lib"

//import "path_pack/chapter/demo2/lib" //注意包的导入方式

func main() {
	name1 := "kevin"
	lib.Hello1(name1) // 直接使用 包名.函数名称来调用函数

}
