package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 定义核酸检测结果的三种状态常量
const (
	YIN    = "阴"
	YANG   = "阳"
	DANGER = "危险" //10混1查出的整管如果是阳，则10个人的状态都置为Danger
)

// 定义试管颜色常量
const (
	PINK  = "粉色" //未经检测的试剂颜色
	RED   = "红色" //阳性
	GREEN = "绿色" //阴性
)

// People 定义核酸样本采样对象结构体 The object of the sample collection = People
type People struct {
	name string
	id   string
	flag string //病毒携带状态
}

// NucleicAcidSample 定义核酸样本结构体,10混1
type NucleicAcidSample struct {
	Peoples [10]*People
	Flag    string //混管检查状态
}

func main() {
	p1 := People{"Zhangsan", "110123198505123856", YIN}
	p2 := People{"Lisi", "110123198305123827", YIN}
	p3 := People{"Langwu", "110123198205123817", YIN}
	p4 := People{"Zhaoliu", "110123198805123457", YIN}
	p5 := People{"Tom", "110123198802123817", YIN}
	p6 := People{"Alice", "110123199505123257", YIN}
	p7 := People{"Lucy", "110123199105123187", YIN}
	p8 := People{"奥特曼", "11012312000512385X", YIN}
	p9 := People{"潘老师", "110123096305123157", YIN}
	p10 := People{"刘备", "110123021105123852", YIN}

	var nas NucleicAcidSample
	nas.Flag = YIN
	nas.Peoples[0] = &p1
	nas.Peoples[1] = &p2
	nas.Peoples[2] = &p3
	nas.Peoples[3] = &p4
	nas.Peoples[4] = &p5
	nas.Peoples[5] = &p6
	nas.Peoples[6] = &p7
	nas.Peoples[7] = &p8
	nas.Peoples[8] = &p9
	nas.Peoples[9] = &p10

	//fmt.Printf("%T,%v\n", nas, nas)
	//for _, v := range nas.Peoples {
	//	fmt.Println(*v)
	//}

	fmt.Println("-----------------------检测之前--------------------")
	for _, v := range nas.Peoples {
		fmt.Println(*v)
	}
	//混管检测
	CheckFor10(&nas)
	fmt.Println("-----------------------混管检测结果--------------------")
	for _, v := range nas.Peoples {
		fmt.Println(*v)
	}

	if nas.Flag == YANG { //如果混管检测为阳，需要单管采样复检或者使用抗原自检
		for _, v := range nas.Peoples {
			CheckForSingle(v)
		}
	}

	fmt.Println("-----------------------单管最终检测结果--------------------")
	for _, v := range nas.Peoples {
		fmt.Println(*v)
	}
	fmt.Println("-----------------------潘老师的最终检测结果--------------------")
	fmt.Println(p9)

}

// CheckFor10 混管检测 10混1
func CheckFor10(nas *NucleicAcidSample) { //注意这个地方的参数一定要传指针才行
	reagentColor := PINK //试剂的初始颜色
	//检测结果以随机数来确定
	randNumb := GetRand()
	//fmt.Println(randNumb)
	switch randNumb {
	case 1:
		reagentColor = RED
	case 2:
		reagentColor = GREEN
	default:
		reagentColor = PINK
	}

	switch reagentColor {
	case RED:
		nas.Flag = YANG
	case GREEN:
		nas.Flag = YIN
	default:
		fmt.Println("未变色，试剂问题检测失败")
	}

	//给10个人置检测结果状态
	for _, v := range nas.Peoples {
		if nas.Flag == YIN {
			v.flag = YIN
		} else {
			v.flag = DANGER
		}
	}
}

// CheckForSingle 单管检测
func CheckForSingle(p *People) {
	reagentColor := PINK //试剂的初始颜色
	randNumb := GetRand()
	//fmt.Println(randNumb)
	switch randNumb {
	case 1:
		reagentColor = RED
	case 2:
		reagentColor = GREEN
	default:
		reagentColor = PINK
	}
	switch reagentColor {
	case RED:
		p.flag = YANG
	case GREEN:
		p.flag = YIN
	default:
		fmt.Println("未变色，试剂问题，检查失败")
	}
}

// AntigenDetection 抗原检测
func AntigenDetection() {

}

func GetRand() int32 {
	//生成随机数 1或2
	min := int32(0) //设置随机数下限
	max := int32(3) //设置随机数上限
	rand.Seed(time.Now().UnixNano())
	num := rand.Int31n(max-min-1) + min + 1
	return num
}
