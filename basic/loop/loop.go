package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)
// convert to binary number
func convertToBin(n int) string {
	result := ""
	if n == 0 {
		result = "0"
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result //反过来＋就不用最后倒转了
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename) //Open返回filename,err两个参数
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	// bufio is an IO package with cache
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),     //101
		convertToBin(13),    //1101
		convertToBin(13333), //1101
		convertToBin(0),     //1101
	)

	printFile("d:/learngo/gocode/basic/basic/abc.txt")
	// the middle of `` is a multiline string
	s := `abc"d"
	kkk
	123
	
	s`
	printFileContents(strings.NewReader(s))

}
