package main

import (
	"fmt"
)

type MyStructs struct {
	x int
	y string
}

func main() {
	fmt.Println("字符串格式化")
	MyNum := 100
	MyStr := "字符串"
	myStruct := MyStructs{1, "jom"}
	// %v可以打印出任何字段的值
	fmt.Printf("MyNum=%v\n", MyNum)    // MyNum=100
	fmt.Printf("MyStr=%v\n", MyStr)    // MyStr=字符串
	fmt.Printf("MyStr=%v\n", myStruct) // MyStr={1 jom}
	// %+v打印出的值将会比较详细
	fmt.Printf("MyNum=%+v\n", MyNum)    // MyNum=100
	fmt.Printf("MyStr=%+v\n", MyStr)    // MyStr=字符串
	fmt.Printf("MyStr=%+v\n", myStruct) // MyStr={x:1 y:jom}
	// %#v打印出的值将会比较详细
	fmt.Printf("MyNum=%#v\n", MyNum)    // MyNum=100
	fmt.Printf("MyStr=%#v\n", MyStr)    // MyStr="字符串"
	fmt.Printf("MyStr=%#v\n", myStruct) // MyStr=main.MyStruct{x:1, y:"jom"}
}
