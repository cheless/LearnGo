package main

import (
	"fmt"
)

func main() {
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "fabulous",
	}
	m2 := make(map[string]int) //m2 == empty map

	var m3 map[string]int //m3 == nil

	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map")
	//遍历是无序的，因为是hashmap
	for k, v := range m {
		fmt.Println(k, v)
	}
	//判断元素
	fmt.Println("Getting values")
	courseName, isok := m["course"]
	fmt.Println(courseName, isok)
	if courseName, isok := m["cause"]; isok {
		fmt.Println(courseName)
	} else {
		fmt.Println("key does not exist")
	}
	//删除元素
	fmt.Println("Deleting values:")
	name, isok := m["name"]
	fmt.Println(name, isok)

	delete(m, "name")
	name, isok = m["name"]
	fmt.Println(name, isok)
}
