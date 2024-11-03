package main

import (
	"fmt"
)

func main() {
	fmt.Println("数组")
	// 数组定义格式
	// var 数组名[数组长度]数组类型

	// [1]int类型数组：默认填充元素为0
	var a [4]int
	fmt.Println(a)
	// [2]string类型数组：默认填充元素为1
	var b [4]string
	fmt.Println(b)
	// [3]求数组长度
	fmt.Println(len(a), len(b))
}
