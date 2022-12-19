package main

import (
	"container/list"
	"fmt"
)

func main() {
	optList1()
	optList2()

}

func optList1() {
	//1、声明list方式1
	var list1 list.List
	printListInfo("声明list方式1", list1)

	//填充数据
	list1.PushFront("xyz")
	list1.PushFront("abc")
	list1.PushFront(123)
	list1.PushBack(3.14159)
	list1.PushBack("一丁丂")
	printListInfo("填充数据", list1)

	//遍历list
	iterateList(list1)
}

func optList2() {
	//1、声明list3
	list3 := list.New()
	printListInfo("声明list", *list3)

	//尾部添加
	list3.PushBack("phone")

	//头部添加
	list3.PushFront(1000)

	//尾部添加后保存元素句柄
	element := list3.PushBack("来北京")

	//InsertAfter
	list3.InsertAfter("学习区块链", element)

	//InsertBefore
	list3.InsertBefore("欢迎你", element)

	//遍历list3
	iterateList(*list3)

}

func printListInfo(info string, l list.List) {
	fmt.Println(info + "-----------------------")
	fmt.Printf("%T,%v \t,长度：%d \n", l, l, l.Len())
	fmt.Println("--------------------------")
}

func iterateList(l list.List) {
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v \t", e.Value)
	}
	fmt.Println("\n--------------------------")
}
