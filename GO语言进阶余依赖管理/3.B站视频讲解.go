package main

import "fmt"
import "time"

func count1(n int, MyStr string) {
	for i := 0; i < n; i++ {
		fmt.Println(MyStr)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	fmt.Println("B")
	go count1(5, "🐏")
	count1(5, "🐂")
	// 运行结果见"go语言并发编程1.png"
}
