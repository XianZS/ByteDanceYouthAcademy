package main

import "fmt"

func deferFunc() int {
	fmt.Println("deferFunc running")
	return 0
}

func returnFunc() int {
	fmt.Println("returnFunc running")
	return 0
}

func allFunc() int {
	defer deferFunc()
	return returnFunc()
}

func main() {
	fmt.Println("defer")
	/*
		[1] defer 关键字的执行时间：
		当前函数生命周期完全结束之后会触发defer栈
		[2] defer 的执行顺序
		程序每次遇到一个 defer ，就会被将这个 defer 压入栈内
		[3] defer 和 return 的执行顺序
		return 的执行顺序高于 defer
	*/
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	fmt.Println("begin :: main")
	fmt.Println("begin :: main")
	/*
		defer
		begin :: main
		begin :: main
		main end2
		main end1
	*/
	allFunc()
}
