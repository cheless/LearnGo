package main

import "fmt"

//值类型数组,传入参数被拷贝（不同于其他语言的默认引用）
func printArray(arr [5]int) {
	arr[0] = 100
	for i, v := range arr { //遍历下标和值
		fmt.Println(i, v)
	}
}

// new try
func main() {
	var arr1 [5]int         //初始值为0
	arr2 := [3]int{1, 3, 5} //局部变量需要初始化
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int //二维数组

	fmt.Println(grid)
	fmt.Println(arr1, arr2, arr3)

	fmt.Println("printArray(arr1)")
	printArray(arr1)
	fmt.Println("printArray(arr3)")
	printArray(arr3)
	fmt.Println("print origin arr1 and arr3")
	fmt.Println(arr1, arr3)
}
