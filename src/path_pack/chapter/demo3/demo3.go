package main

import (
	lib3 "path_pack/chapter/demo3/lib"
) //注意包的导入方式,最后一个是目录名而不是包名

func main() {
	name := "kevin"
	lib3.Hello(name) // 直接使用 包名.函数名称来调用函数

}
