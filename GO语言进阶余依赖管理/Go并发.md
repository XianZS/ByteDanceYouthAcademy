# 协程

## 1、概念

> 线程属于内核态，比较消耗资源，栈MB级别

> 协程：用户态，轻量级线程，栈KB级别

## 2、语法

```go
go func(传入参数A 传入参数A类型,传入参数B 传入参数B类型,……){
	函数体
}(传入参数A 传入参数A类型,传入参数B 传入参数B类型,……)
```

## 3、协程之间的通信

### (1) 基本逻辑关系

通过通信共享内容，而不是通过共享内存来实现通信

### (2) Channel类型

make(chan 元素类型, \[ 缓冲大小 ] )

* 无缓冲通道

    &emsp;&emsp;使用无缓冲通道发送数据时，需要保证发送方和接收方都同时准备好，这样才能完成通信。

    &emsp;&emsp;如果两个goroutine没有同时准备好，通道会导致先执行发送或接收操作的goroutine阻塞等待。

* 有缓冲通道

    &emsp;&emsp;有缓冲通道，指通道可以保存多个值. 

    &emsp;&emsp;如果给定了一个缓冲区容量，那么通道就是异步的，只要缓冲区有未使用空间用于发送数据，或还包含可以接收的数据，那么其通信就会无阻塞地进行

### (3)理解

> Channel（管道）可以被认为是协程之间通信的管道。

> 与水流从管道的一端流向另一端一样，数据可以从信道的一端发送并在另一端接收。

### (4)示例代码

* 生产者：A协程 
    > 生产者只需要顾及生产，所以使用不带缓冲区的Channel来模拟

* 消费者：B协程
    > 消费者需要进行一些复杂操作，所以使用带缓冲区的Channel来模拟

```go
package main

import "fmt"

func CalSquare() {
	src := make(chan int)     // 定义无缓冲队列
	dest := make(chan int, 3) // 有缓冲队列
	// A协程的功能
	go func() {
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
		// 可能存在的复杂操作
		fmt.Println(i/2, "+", i/2, "=", i)
	}
}

func main() {
	fmt.Println("Channel")
	CalSquare()
}
```

> defer语句的用法：
> 
> 【1】defer语句会将函数推入到一个列表中。
> 
> 【2】同时列表中的函数会在return语句执行后被调用。
> 
> 【3】defer常常会被用来简化资源清理释放之类的操作。