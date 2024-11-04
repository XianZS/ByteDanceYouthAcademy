package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("字符串和数字之间的转换")
	some1 := "123"
	newSome1, err1 := strconv.Atoi(some1)
	fmt.Println(err1, ":", some1, "转换为数字之后的结果为:", newSome1)
	// <nil> : 123 转换为数字之后的结果为: 123
	some2 := "AAA"
	newSome2, err2 := strconv.Atoi(some2)
	fmt.Println(err2, ":", some2, "转换为数字之后的结果为:", newSome2)
	// strconv.Atoi: parsing "AAA": invalid syntax : AAA 转换为数字之后的结果为: 0
}
