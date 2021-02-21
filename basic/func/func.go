package main

import (
	"fmt"
	"math"
	"reflect"
)

//返回单个值
func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		r, _ := div(a, b)
		return r, nil
	default:
		return 0, fmt.Errorf(
			"unsupported operation: %s", op)
	}
}

//返回多个值
func div(a, b int) (q, r int) {
	return a / b, a % b
}

//函数式编程！
//参数为一个返回int的函数op和两个int变量a,b
func apply(op func(int, int) int, a, b int) int {
	opName := reflect.ValueOf(op).Pointer()
	fmt.Printf("Calling function %s with args"+
		"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(13, 4, "x"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	fmt.Println(q, r)

	fmt.Println(apply(
		//参数为匿名函数
		func(a int, b int) int {
			return int(math.Pow(
				float64(a), float64(b)))
		}, 3, 4))

	fmt.Println(sum(1, 2, 3, 4))
}
