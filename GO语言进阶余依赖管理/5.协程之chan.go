package main

import (
	"fmt"
	"time"
)

func Count(animals string, c int, myChan chan string) {
	for i := 0; i < c; i++ {
		myChan <- animals
		time.Sleep(time.Millisecond * 500)
	}
	// 进程对象被关闭
	close(myChan)
}

func main() {
	//wg.Add(2) // 添加协程数

	// 针对4.go文件中出现的统计动物数量问题，
	c := make(chan string)
	/*
		一个chan只能发送和接收一种数据类型
		往chan中发送/接收数据时，会阻塞代码
	*/
	go Count("🐏", 5, c)
	/*
		for {
			message := <-c
			fmt.Println(message)
			// fatal error: all goroutines are asleep - deadlock!
		}
	*/
	for animal := range c {
		fmt.Println(animal)
	}
}
