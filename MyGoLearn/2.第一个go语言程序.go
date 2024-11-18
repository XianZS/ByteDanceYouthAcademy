package main

// 程序包含的包名

import (
	"fmt"
	"time"
)

// go语言要求左花括号和当前函数名属于同一行
func main() {
	fmt.Println("This is my first go language func")
	time.Sleep(1 * time.Second)
}
