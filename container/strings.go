package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "yes我是富贵猪"
	fmt.Println(s)
	//转为byte切片发现是unit8
	//utf-8是可变长的，每个中文字符占3位，每个英文字符占1位
	for i, b := range []byte(s) {
		fmt.Printf("%d:%X ", i, b) //%X为16进制数
	}
	fmt.Println()
	//为适应多种语言，string的字符类型为rune，一个rune相当于int32,（4字节）
	//rune相当于go的char
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	//获得字符数量
	fmt.Println("Rune count:",
		utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		//DecodeRune解码UTF-8编码序列的第一个UTF8编码，返回该码值和长度
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:] //裁剪字符串
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()
	fmt.Println()

}
