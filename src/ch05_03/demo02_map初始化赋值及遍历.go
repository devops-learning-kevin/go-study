package main

import "fmt"

func main() {
	//1、声明时同时初始化
	var country = map[string]string{
		"China":  "Beijing",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
		"France": "Paris",
		"Italy":  "Rome",
	}
	fmt.Println(country)

	//短变量声明初始化方式
	rating := map[string]float64{"c": 5, "Go": 4.5, "Python": 4.5, "C++": 3}
	fmt.Println(rating)

	//2、创建map后再赋值
	countryMap := make(map[string]string)
	countryMap["China"] = "Beijing"
	countryMap["China"] = "Beijing"
	countryMap["Japan"] = "Tokyo"
	countryMap["India"] = "New Delhi"
	countryMap["France"] = "Paris"
	countryMap["Italy"] = "Rome"

	//3、遍历map 无序
	//（1）、key、value都遍历
	for k, v := range countryMap {
		fmt.Println("国家", k, "首都", v)
	}
	fmt.Println("-----------------")

	//（2）、只展示value
	for _, v := range countryMap {
		fmt.Println("国家", "首都", v)
	}
	fmt.Println("-----------------")

	//（3）、只展示key
	for k := range countryMap {
		fmt.Println("国家", k, "首都", countryMap[k])
	}
	fmt.Println("-----------------")

	//4、查看元素是否在map中存在
	value, ok := countryMap["England"]
	fmt.Printf("%q\n", value)
	fmt.Printf("%T,%v\n", ok, ok)
	if ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都信息未检索到")
	}

	//或者
	if value, ok := countryMap["USA"]; ok {
		fmt.Println("首都：", value)
	} else {
		fmt.Println("首都信息未检索到")

	}

}
