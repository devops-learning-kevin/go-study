package main

import "fmt"

//switch语句，只执行一个分支
func main() {
	ScoreGrade()

}

func ScoreGrade() {
	score := 86
	grade := ""

	switch { //swith后面的表达式为空的话，就是ture
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	switch { //swith后面的表达式为空的话，就是ture
	case score >= 90:
		grade = "A"
	case score >= 80:
		grade = "B"
	case score >= 70:
		grade = "C"
	case score >= 60:
		grade = "D"
	default:
		grade = "E"
	}
	fmt.Printf("等级是：%s\n", grade)
	fmt.Printf("评价是：")

	switch grade {
	case "A":
		fmt.Println("优秀")
	case "B":
		fmt.Println("良好")
	case "C":
		fmt.Println("中等")
	case "D":
		fmt.Println("及格")
	default:
		fmt.Println("差")
	}

}
