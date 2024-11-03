package main

import (
	"fmt"
)

func main() {
	// 不需要自动推导，来进行赋值
	var a = "some"
	var b = 's'
	var c = 32
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	// 自动推导来进行赋值
	d := 2
	fmt.Println("d:", d)
	// 定义常量
	const s string = "const string"
	fmt.Println("s:", s)
}
