package main

import "fmt"

func main() {
	//初始薪资
	salary := map[string]float64{
		"Steven": 15000,
		"John":   20000,
		"Daniel": 5000,
	}
	fmt.Println("初始薪资", salary)

	//调薪后
	newSalary := salary
	newSalary["Daniel"] = 8000
	fmt.Println("调薪后：", newSalary)
	fmt.Println("原始薪资salary是都受影响", salary)

	//定期涨薪
	changeSalary(salary)
	fmt.Println("原始薪资salary是都受影响", salary)

}

//定期涨薪
func changeSalary(m map[string]float64) {
	for k := range m {
		m[k] *= 1.1
	}
}
