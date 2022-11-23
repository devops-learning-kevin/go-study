package main

import "fmt"

func main() {
	GetDaysByMonth()
}

//switch语句在go中和java是不同的，java需要再case分支结束时，加break语句，而go不用加break。
// case分支只能选择走一个。

//判断某年某月有多少天
func GetDaysByMonth() {
	//第一局部变量年月日
	year := 2008
	month := 2
	days := 0

	switch month {
	case 1, 3, 5, 7, 8, 10, 12: //go的switch语句，case后面可以有多个结果
		days = 31
	case 4, 6, 9, 11:
		days = 30
	case 2:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			days = 29
		} else {
			days = 28
		}
	default:
		days = -1
	}

	fmt.Printf("%d年%d月有%d天", year, month, days)
}
