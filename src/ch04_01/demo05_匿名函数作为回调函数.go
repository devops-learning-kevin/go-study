package main

import (
	"fmt"
	"math"
)

type myFuncs func(float64) string

func main() {
	//定义切片，对其中的数据元素进行求平方根和平方的运算
	arr := []float64{1, 9, 25, 36, 250}

	//遍历切片，对每个元素进行处理运算
	temp := FilterSlice(arr, func(data float64) string {
		return fmt.Sprintf("%.0f", math.Sqrt(data))
	})
	fmt.Println(temp)

	temp = FilterSlice(arr, func(data float64) string {
		return fmt.Sprintf("%.0f", math.Pow(data, 2.0))
	})
	fmt.Println(temp)

}

func FilterSlice(arr []float64, f myFuncs) []string {
	var result []string
	for _, v := range arr {
		result = append(result, f(v))
	}
	return result
}
