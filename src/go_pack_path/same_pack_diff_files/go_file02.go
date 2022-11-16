package same_pack_diff_files

import "fmt"

//私有变量，首字母小写，同一个包的文件可以直接调用，外包不能调用
var file02_var01 int32 = 201

//公有变量，首字母大写，同一个包的文件可以直接调用，跨包可以使用包名.变量名调用
var File02_var02 int32 = 202

//私有常量，首字母小写，只有同一个包的文件才可以调用，并且直接调用，跨包不能调用
const file02_const01 = 2001

//公共常量，首字母大写(通常是全部大写),同一个包的文件，可以直接调用，跨包调用需要加上包名.
const File02_const02 = 2002
const FILE02_CONST03 = 2003

//私有函数，只能本包调用，直接调用,不能跨包调用
func file02_func01() {
	fmt.Println("This is func file02_func01.")
}

//公有函数，本包直接调用，跨包调用加上包名.
func File02_func02() {
	fmt.Println("This is func File02_func02.")
}
