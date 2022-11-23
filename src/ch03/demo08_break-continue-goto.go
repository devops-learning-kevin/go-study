package main

import "fmt"

func main() {
	testBreak()

	//break和continue的区别
	breakContinue()

	//输出1-50之间不包含4的数字（个位和十位）continue实现
	eludeFour()

	//输出1-50之间不包含4的数字（个位和十位）goto实现
	eludeFourGoto()

	//输出1-100之间的素数 (借助togo跳转)
	printPrimeNumber()
}

/*
循环控制语句：break,continue
break,continue的区别：
break将无条件的跳出并结束当前的循环，然后执行循环后的语句
continue语句是跳过当前的某一次循环，而开始执行下一次循环
break和continue只对当前层的循环有效，对外部循环无效
*/

func testBreak() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\n外部部循环i =%d,", i)
		for j := 0; j < 10; j++ {
			if j == 5 {
				break //直接跳出当前循环
				//continue //跳过本次循环，继续下一次循环
			}
			fmt.Printf("\n内部循环i =%d,j=%d", i, j)
		}
	}
}

func breakContinue() {
	fmt.Println("\n1、break 终止循环")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i)
	}
	//结果是01234

	fmt.Println("\n2、continue跳过某次循环")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
	}
	//结果是012346789
}

//输出1-50之间不包含4的数字（个位和十位）
func eludeFour() {
	fmt.Println("\n输出1-50之间不包含4的数字")
	num := 0
	for num < 50 {
		num++
		if (num%10 == 4) || (num/10%10) == 4 {
			continue
		}
		fmt.Printf("%d ", num)
	}

}

//输出1-50之间不包含4的数字（个位和十位）(goto实现)
func eludeFourGoto() {
	fmt.Println("\n输出1-50之间不包含4的数字")
	num := 0
LOOP:
	for num < 50 {
		num++
		if (num%10 == 4) || (num/10%10 == 4) {
			goto LOOP
		}
		fmt.Printf("%d ", num)
	}
}

//输出1-100之间的素数 (借助togo跳转)
func printPrimeNumber() {
	fmt.Println("\n输出1-100之间的素数 (借助togo跳转)")
	num := 0
LOOP:
	for num < 100 {
		num++
		for i := 2; i < num; i++ {
			if num%i == 0 {
				//break  //输出的是1-100
				//continue //输出的是1-100
				//fmt.Printf("\n%d 不是素数", num)
				goto LOOP
			}
		}
		fmt.Printf("%d ", num)
	}

}
