package main

import (
	"fmt"
	"same_pack_diff_files/same_pack_diff_files"
)

func main() {

	//调用 go_file01中的全局公有常量
	fmt.Println("go_file01中的全局常量：File01_const02 = ", same_pack_diff_files.File01_const02)
	fmt.Println("go_file01中的全局常量：FILE02_CONST03 = ", same_pack_diff_files.FILE01_CONST03)

	//调用 go_file01中的全局公有变量
	fmt.Println("go_file01中的全局变量：File01_var02 = ", same_pack_diff_files.File01_var02)

	//调用 go_file01中的函数公有File01_func02
	same_pack_diff_files.File01_func02()

	//调用 go_file02中的公有函数File02_func02
	same_pack_diff_files.File02_func02()

	//调用 go_file01中的公有函数File01_call_file02
	same_pack_diff_files.File01_call_file02()
}
