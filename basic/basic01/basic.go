package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//包内变量，非全局变量
var aa = 3
var ss = "kkk"
var bb = true

//变量自动初始化
func variableZeroValue() {
	var (
		a int
		s string
	)
	fmt.Printf("%d %q\n", a, s)
}

//手动初始化
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

//变量自动判断类型
func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

//局部变量用:=快速定义
func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

//欧拉公式计算
func euler() {
	c := cmplx.Pow(math.E, 1i*math.Pi) + 1 //结果应该为0，但是浮点数计算精读有限
	//	=complex.Exp(1i * math.PI) + 1
	fmt.Printf("%.1f\n", c) //用.f控制精读
}

//勾股定理计算,需注意的是go中不能进行自动类型转换
func triangle() {
	var (
		a, b = 3.0, 4.0
		c    float64
	)
	c = math.Sqrt(a*a + b*b)
	/*
		= var(
			a, b = 3, 4
			c	int
		)
		c = (int)(math.Sqrt(float64(a*a + b*b)))
	*/
	fmt.Println(c)
}

//常量
const filename = "variable.go"

func consts() {
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

//枚举类型
func enums() {
	const (
		golang = iota //iota为常量自增计数器
		_             //1
		cpp           //2
		java
		python
	)
	//iota的扩展用法
	const (
		b  = 1 << (10 * iota)
		kb //相当于	1 << (10 * 1)
		mb
		gb
		tb
		pb
	)

	fmt.Println(golang, cpp, java, python)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, ss, bb)
	euler()
	triangle()
	consts()
	enums()
}
