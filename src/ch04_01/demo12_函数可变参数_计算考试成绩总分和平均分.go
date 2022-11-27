package main

import "fmt"

func main() {
	//1、传n个成绩的用法
	sum, avg, count, user := GetScore("kevin", 78, 85, 69, 98, 88, 92)
	fmt.Printf("%s同学%d门课的总成绩是：%.2f,平均成绩是：%.2f,", user, count, sum, avg)

	fmt.Println()

	//2、传切片的用户
	scores := []float64{88, 67, 94, 86, 90, 79, 99, 69}
	sum, avg, count, user = GetScore("xiaopeng", scores...)
	fmt.Printf("%s同学%d门课的总成绩是：%.2f,平均成绩是：%.2f,", user, count, sum, avg)

}

func GetScore(name string, scores ...float64) (sum, avg float64, count int, user string) {

	for _, v := range scores {
		sum += v
		count++
	}
	user = name
	avg = sum / float64(count)
	return
}
