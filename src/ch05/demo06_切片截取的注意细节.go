package main

import "fmt"

func main() {
	sliceCap()
}

func sliceCap() {
	arr := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	fmt.Println("cap(arr) = ", cap(arr), arr) //[a b c d e f h i j k]

	//截取数组，形成切片
	s1 := arr[2:8]
	fmt.Printf("%T \n", s1)              //[]string 切片类型，不是数组类型
	fmt.Println("cap(s1)=", cap(s1), s1) // 9  [c d e f h i]

	s2 := arr[4:7]
	fmt.Println("cap(s2)=", cap(s2), s2) // 7 [e f g]

	//截取切片，形成新切片
	s3 := s1[3:9]
	//s3 := s1[3:10]  //panic: runtime error: slice bounds out of range [:10] with capacity 9
	fmt.Println("截取s1[3:9]后形成的s3=", s3) //[f g h i j k]

	//截取切片，形成切片
	s4 := s2[4:7]
	fmt.Println("截取s2[4:7]后形成的s4=", s4) //[i j k]

	//切片是引用类型，是指向底层数组的地址
	s4[0] = "x"
	fmt.Println(arr, s1, s2, s3, s4) //发现数组arr[8] 变了

}
