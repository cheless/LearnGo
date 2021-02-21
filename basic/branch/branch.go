package main

import (
	"fmt"
	"io/ioutil"
)

//switch写法,go中省略break，如不需要break则声明
func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf(
			"Wrong score: %d", score))
	}
	return g
}
func main() {
	//这种相对路径一定要保证工作路径一致，不仅仅是运行路径
	const filename = "abc.txt"
	//go语言的独特if写法
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(88),
		grade(99),
		grade(100),
		grade(101), //panic终端执行，判断出错
	)
}
