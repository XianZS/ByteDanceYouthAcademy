package main

import (
	"fmt"
)

func main() {
	// GoLand只存在for一种循环
	fmt.Println("for循环")
	var i = 1
	for i < 3 {
		fmt.Println("i : ", i)
		i++
	}
	// continue和break依旧可以用
}
