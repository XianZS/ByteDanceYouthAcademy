package main

import (
	"fmt"
	"time"
)

func use() {
	fmt.Println()
	time.Sleep(time.Millisecond * 500)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for {
			c1 <- "🐏"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "🐂"
			time.Sleep(time.Millisecond * 2000)
		}
	}()
	/*
		[1]错误方式
		如果这里使用for直接输出的话，会存在如下情况：
		for{
			fmt.Println(<-c1)
			fmt.Println(<-c2)
		}
		会发现这个for循环运行的代码逻辑并不是我们想象中的那样：“存在时间时间间隔的输出”，具体原因见“6-Chan的管理”文件夹下的照片
	*/
	// [2]正确方式，使用select语句
	for {
		select {
		case msg := <-c1:
			fmt.Print(msg + " ")
		case msg := <-c2:
			fmt.Print(msg + " ")
		}
	}
	/*
		正确运行结果如下：
		🐂 🐏 🐏 🐏 🐏 🐂 🐏 🐏 🐏 🐏 🐂 🐏 🐏 🐏 🐏 🐂 🐏 🐏 🐏 🐏 🐂 🐏 🐏 🐏 🐏 🐂 🐏 🐏 🐏 🐏 🐂 🐏
	*/
}

/*
	[1]select-case结构语句：
	跟 switch-case 相比，select-case 用法比较单一，它仅能用于 信道/通道 的相关操作。
	[2]语法
	```
	select {
	    case 表达式1:
	        <code>
	    case 表达式2:
	        <code>
	  default:
	      <code>
	}
	```
	[3]特性
	[3.1]避免造成死锁
		select 在执行过程中，必须命中其中的某一分支。
		如果在遍历完所有的 case 后，若没有命中
		（命中：也许这样描述不太准确，我本意是想说可以执行信道的操作语句）
		任何一个 case 表达式，就会进入 default 里的代码分支。
		但如果你没有写 default 分支，select 就会阻塞，直到有某个 case 可以命中，
		而如果一直没有命中，select 就会抛出 deadlock 的错误，就像下面这样子。
	[3.2]随机性
	之前学过 switch 的时候，知道了 switch 里的 case 是顺序执行的，但在 select 里却不是。
*/
