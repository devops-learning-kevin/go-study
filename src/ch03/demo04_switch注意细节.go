package main

import "fmt"

/*
(二)、switch语句中的注意细节
1、switch语句执行的过程自上而下，直到找到case匹配项，匹配项中无需使用break，因为Go语言中的switch默认给每个case自带break，因此匹配成功后不会向下执行其他的case分支，而是跳出整个switch。
2、变量var1可以是任何类型，而val1和val2则可以是同类型的任意值。类型不被局限干常量或整数，但必须是相
同类型或最终结果为相同类型的表达式。
3、case后的值不能重复。
4、可以同时测试多个符合条件的值，也就是说case后可以有多个值，这些值之间使用逗号分割，例如:case val1, val2,val3。
5、Go语言中switch后的表达式可以省略，那么默认是switch true。
6、Go语言中的switchcase因为自带break，所以匹配某个case后不会自动向下执行其他case，如需贯通后续的 case，可以添加fallthrough(中文含义是:贯穿)，强制执行后面的case分支。
7、fallthrough必须放在case分支的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误(fallthrough statement out of place，含义是fallthrough不在合适的位置)。
*/

func main() {
	eval()

}

func eval() {
	num1, num2, result := 18, 12, 0
	operator := "%"

	switch operator {
	case "+":
		result = num1 + num2
		fmt.Println("执行+")
		//fallthrough
	case "-":
		result = num1 - num2
		fmt.Println("执行-")
		fallthrough
	case "*":
		result = num1 * num2
		fmt.Println("执行*")
		fallthrough
	case "/":
		result = num1 / num2
		fmt.Println("执行/")
		//fallthrough    没有fallthrough，后面的case就不会贯穿执行了
	case "%":
		result = num1 % num2
		fmt.Println("执行%")
		fallthrough
	default:
		fmt.Println("啥都不执行")
		result = -1

	}
	fmt.Println(result)
}
