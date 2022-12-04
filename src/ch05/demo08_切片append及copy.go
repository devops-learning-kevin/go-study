package main

import "fmt"

func main() {
	fmt.Println("1、---------------切片append------------------------")
	var numbers []int
	printSlice("numbers:", numbers) //[]

	if numbers == nil {
		fmt.Println("numbers==nil", len(numbers))
	}

	//append一个元素
	numbers = append(numbers, 0)
	printSlice("numbers:", numbers) //[0]

	//append多个元素
	numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7)
	printSlice("numbers:", numbers) //value=[0 1 2 3 4 5 6 7 8]

	//append一个切片
	s1 := []int{100, 200, 300, 400, 500, 600, 700}
	numbers = append(numbers, s1...)
	printSlice("numbers:", numbers) //value=[0 1 2 3 4 5 6 7 8 100 200 300 400 500 600 700]

	fmt.Println("2、--------------切片删除-------------------------")
	//删除第一元素
	numbers = numbers[1:]
	printSlice("numbers:", numbers) //value=[1 2 3 4 5 6 7 100 200 300 400 500 600 700]

	//删除最后一个元素
	numbers = numbers[:len(numbers)-1]
	printSlice("numbers:", numbers) //value=[1 2 3 4 5 6 7 100 200 300 400 500 600]

	//删除中间一个元素
	a := int(len(numbers) / 2)
	fmt.Println("中间下标:", a)
	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlice("numbers:", numbers) //value=[1 2 3 4 5 6 100 200 300 400 500 600]

	fmt.Println("3、--------------切片copy-------------------------")
	//创建目标切片
	numbers1 := make([]int, len(numbers), cap(numbers)*2)

	//将numbers中的元素拷贝到numbers1
	count := copy(numbers1, numbers)
	fmt.Println("拷贝的个数：", count)
	printSlice("numbers1", numbers1)

	//验证拷贝的两个切片是否有关联
	numbers[0] = 99
	numbers1[len(numbers1)-1] = 111

	printSlice("numbers", numbers)
	printSlice("numbers1", numbers1)

}

func printSlice(name string, x []int) {
	fmt.Print(name, "\t")
	fmt.Printf("地址：%p \t  len=%d \t  cap=%d \t value=%v \n", x, len(x), cap(x), x)
}
