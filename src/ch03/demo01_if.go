package main

import "fmt"

func main() {
	EvenOdd()
	ScoreMark()
	ScoreMark2()
	ScoreMark3()

}

//判断数值的奇偶性
func EvenOdd() {
	num := 30
	if num%2 == 0 {
		fmt.Println(num, "偶数")
	} else {
		fmt.Println(num, "奇数")
	}
}

//判断学生成绩
func ScoreMark() {
	score := 78
	if score >= 90 {
		fmt.Println("优秀")
	}

	if score >= 80 && score < 90 {
		fmt.Println("良好")
	}

	if score >= 70 && score < 80 {
		fmt.Println("中等")
	}

	if score >= 60 && score < 70 {
		fmt.Println("及格")
	}

	if score < 60 {
		fmt.Println("不及格")
	}

}

// if...else if...else
func ScoreMark2() {
	score := 88
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 70 {
		fmt.Println("中等")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Printf("不及格")
	}
}

// if嵌套
func ScoreMark3() {
	score := 98
	if score >= 60 {
		if score >= 70 {
			if score >= 80 {
				if score >= 90 {
					fmt.Println("优秀")
				} else {
					fmt.Printf("良好")
				}
			} else {
				fmt.Printf("中等")
			}
		} else {
			fmt.Printf("及格")
		}
	} else {
		fmt.Printf("不及格")
	}
}
