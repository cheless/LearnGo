package main

import "fmt"

//用于观察append操作时len和cap的变化
func printSlice(s []int) {
	fmt.Printf("%v, len = %d, cap = %d\n",
		s, len(s), cap(s))
}

func main() {
	fmt.Println("Creating slice")
	//nil是预定义的标识符，代表指针、通道、函数、接口、映射或切片的零值,并不是GO 的关键字之一
	var s []int //定义slice,此时slice为nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	//初始化定义
	s1 := []int{2, 4, 6, 8}
	printSlice(s1)
	//不初始化值但是初始len或者cap
	s2 := make([]int, 16)
	s3 := make([]int, 10, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1)
	printSlice(s2)
	//用下标删除特定元素
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) //...表示切片依次使用
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}
