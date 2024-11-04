package main

import "fmt"

func CalSquare() {
	src := make(chan int)     // 定义无缓冲队列
	dest := make(chan int, 3) // 有缓冲队列
	// A协程的功能
	go func() {
		//defer语句会将函数推入到一个列表中。同时列表中的函数会在return语句执行后被调用。defer常常会被用来简化资源清理释放之类的操作。
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	// B协程的功能
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i + i
		}
	}()
	// 打印输出
	for i := range dest {
		fmt.Println(i/2, "+", i/2, "=", i)
	}
}

func main() {
	fmt.Println("Channel")
	CalSquare()
}
