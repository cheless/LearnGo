package main

import "fmt"

func lengthOfNonRepeatingSubstr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	//每次找到一个start开头，i结尾的不重复子串
	for i, ch := range []rune(s) { //i为下标，ch为字符,此处[]byte(s)有点不清楚
		//如果从左往右扫描过程中该字符在start或之后出现过，则要减去子串的一个子串，以计算新的子串
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		//此时start要么是0要么是上一个重复字符+1
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(
		lengthOfNonRepeatingSubstr("abcabcd"))
	fmt.Println(
		lengthOfNonRepeatingSubstr("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}

/*
ds:
start:当前含有最长不重复子串的开始

*/
