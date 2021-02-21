package main

import "fmt"

func findMaxSubstr(str string) string {
	m := make(map[string]int)
	start := 0
	begin := 0
	end := len(str) - 1
	for start != len(str)-1 {
		//找出以start开头的最长子字符串,末尾标记为end
		begin = start
		for i := begin + 1; begin != end; {
			//如果找到了末尾则跳转新一轮
			if i == end {
				begin++
				continue
			}

			if str[i] == str[begin] {
				end = i
				begin++
			} else {
				i++
			}
		}
		str1 := str[start:end]
		length := end - start
		m[str1] = length
		start++
		end = len(str) - 1
	}
	max := 0
	maxstr := ""
	for k, v := range m {
		if max < v {
			max = v
			maxstr = k
		}
	}
	return maxstr
}

func main() {
	str := "abcabcdab"
	fmt.Println("The biggest str is:", findMaxSubstr(str))
}

/*
ds:
start:标记以当前字符开头的最大子字符串
begin:标记以当前字符至end之前的重复字符查找起始字符
end:当前趟遍历的最后一个字符
al:
1. 从第一个字符到最后一个，每一个字符依次充当start
2. 找出以每个start开头的最大子字符串，存入map
3. 将map中的最大字符串输出
*/
