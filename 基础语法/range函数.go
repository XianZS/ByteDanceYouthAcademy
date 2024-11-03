package main

import (
	"fmt"
)

func main() {
	MyArray := make([]int, 10)
	MyDict := map[string]int{"Jom": 11, "kom": 22, "lom": 33}
	// 遍历数组
	for i, num := range MyArray {
		MyArray[i] = i + 1 + num
	}
	// 遍历字典
	for k, v := range MyDict {
		fmt.Println(k, ":", v)
	}
	fmt.Println(MyArray)
	fmt.Println(MyDict)
}
