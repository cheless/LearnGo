package main

import "fmt"

func updateSlice(s []int) { //[]中没有数表示的才是slice，有数表示array
	s[0] = 100
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//slice基本操作
	s1 := arr[2:6] //s1 = [100, 3, 4, 5]
	s2 := s1[3:5]  //s2 = ?
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n",
		s1, len(s1), cap(s2))
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n",
		s2, len(s2), cap(s2))
	/*
		slice有三个参数：ptr len cap
		ptr指向了slice启示地址，所以reslice不能向前扩展
		len规定了当前的slice范围
		cap规定了reslice的范围
	*/
	//slice append
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	//s3和s4,s5的区别是什么？
	fmt.Println("s3, s4, s5 = ", s3, s4, s5)
	fmt.Println("arr = ", arr)
	//添加元素时如果超越当前slice的cap，系统将分配更大的底层数组
}
