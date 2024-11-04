package main

import (
	"fmt"
)

func main() {
	fmt.Println("错误处理")
	a := 1
	b := 0
	c := a / b
	fmt.Println(c)
}
