package main

import "fmt"

func main() {
	// 1、打印矩形
	printRectangle(10)

	//2、打印左下角的直角三角形
	printRightTriangleLB(10)

	//3、打印左上角的直角三角形
	printRightTriangleLT(10)

	//4、打印右上角的直角三角形
	printRightTriangleRT(10)

	//5、打印右下角的直角三角形
	printRightTriangleRB(10)

	//6、打印等腰三角形
	printEqualTriangle(10)

	//7、打印99乘法表
	multiple99()
}

// 1、打印矩形
func printRectangle(lines int) {
	fmt.Println("\n打印矩形")
	for i := 1; i <= lines; i++ {
		for j := 1; j <= lines; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//2、打印左下角的直角三角形
func printRightTriangleLB(lines int) {
	fmt.Println("打印左下角直角三角形")
	for i := 1; i <= lines; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//3、打印左上角的直角三角形
func printRightTriangleLT(lines int) {
	fmt.Println("打印左上角直角三角形")
	for i := 1; i <= lines; i++ {
		for j := lines; j >= i; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//4、打印右上角的直角三角形
func printRightTriangleRT(lines int) {
	fmt.Println("打印右上角直角三角形")
	for i := 1; i <= lines; i++ {
		for m := 1; m < i; m++ {
			fmt.Print("  ")
		}
		for j := lines; j >= i; j-- {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//5、打印右下角的直角三角形
func printRightTriangleRB(lines int) {
	fmt.Println("打印右下角直角三角形")
	for i := 1; i <= lines; i++ {
		for m := lines; m > i; m-- {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

//6、打印等腰三角形
func printEqualTriangle(lines int) {
	fmt.Println("打印等腰三角形")
	for i := 1; i <= lines; i++ {
		//打印空格
		for m := lines; m >= i; m-- {
			fmt.Print("  ")
		}
		//打印星号
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("* ")
		}
		fmt.Println()

	}
}

//7、打印99乘法表
func multiple99() {
	fmt.Println("打印99乘法表")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d x %d = %d  ", j, i, i*j)
		}
		fmt.Println()
	}

}
