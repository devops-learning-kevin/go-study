package same_pack_diff_files

import "fmt"

//私有变量，首字母小写，同一个包的文件可以直接调用，外包不能调用
var file01_var01 int32 = 101

//公有变量，首字母大写，同一个包的文件可以直接调用，跨包可以使用包名.变量名调用
var File01_var02 int32 = 102

//私有常量，首字母小写，只有同一个包的文件才可以调用，并且直接调用，跨包不能调用
const file01_const01 = 1001

//公共常量，首字母大写(通常是全部大写),同一个包的文件，可以直接调用，跨包调用需要加上包名.
const File01_const02 = 1002
const FILE01_CONST03 = 1003

//私有函数，只能本包调用，直接调用,不能跨包调用
func file01_func01() {
	fmt.Println("This is func file01_func01.")
}

//公有函数，本包直接调用，跨包调用加上包名.
func File01_func02() {
	fmt.Println("This is func File01_func02.")
}

// 同包，不同文件调用示例
func File01_call_file02() {
	fmt.Println("file01_call_file02 start")
	fmt.Println("file02_var01 = ", file02_var01)
	fmt.Println("File02_var02 = ", File02_var02)
	fmt.Println("file02_const01 = ", file02_const01)
	fmt.Println("File02_const02 = ", File02_const02)
	fmt.Println("FILE02_CONST03 = ", FILE02_CONST03)
	file02_func01()
	File02_func02()
	fmt.Println("file01_call_file02 end")
}
