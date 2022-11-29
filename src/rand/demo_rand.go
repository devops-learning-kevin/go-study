package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//生成随机数
	min := int32(0) //设置随机数下限
	max := int32(4) //设置随机数上限

	rand.Seed(time.Now().UnixNano())
	//创建100个随机数
	for i := 0; i < 1000; i++ {
		//不包含上下限
		num := rand.Int31n(max-min-1) + min + 1
		//包含上下限
		// num := rand.Int31n(max-min) + min
		fmt.Print(num, ",")
		if num == min {
			fmt.Println("随机数等于下限---------", num)
		}
		if num == max {
			fmt.Println("随机数等于上限---------", num)
		}
		//脱离范围
		if num < min || num > max {
			fmt.Println("随机数脱离范围区间-----", num)
		}
	}

}
