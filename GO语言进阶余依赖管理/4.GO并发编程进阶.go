package main

import (
	"fmt"
	"sync"
)
import "time"

func count2(n int, MyStr string) {
	for i := 0; i < n; i++ {
		fmt.Print(MyStr + " ")
		time.Sleep(time.Millisecond * 500)
	}
	fmt.Println()
}

func MyMain1() {
	go count2(5, "🐏")
	go count2(5, "🐂")
	/*
		如果将代码“count2(5, "🐂")”改变为“go count2(5, "🐂")”
		再次运行代码，会发现没有任何输出
		*** 因为当主函数main()函数运行完成后，这个程序就结束了，不管存在多少个go协程，都会被杀死 ***
	*/
	// 针对main运行结束，作业结束，go协程被杀死的情况，可以sleep主函数
	time.Sleep(time.Second * 3)
	// 但是这种写法不现实，因为在大多数时候，我们并不知道go协程还需要多长时间才能结束
	// 所以需要调用wg.Add(协程数量)
}

func MyMain2() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		count2(5, "🐏")
		wg.Done()
	}()
	go func() {
		count2(5, "🐂")
		wg.Done()
	}()
	wg.Wait()
}

func MyMain3() {
	var wg sync.WaitGroup
	wg.Add(2)

}

func main() {
	// 针对MyMain1()函数中，主函数main()作业结束，go协程被杀死的情况，可以sleep主函数
	fmt.Println("MyMain1()函数")
	MyMain1()
	/*
		但是这种写法不现实，因为在大多数时候，我们并不知道go协程还需要多长时间才能结束
		所以我们使用wg.Add(协程数量)，来进行相关操作
	*/
	fmt.Println("MyMain2()函数")
	MyMain2()
	/*
		在使用wg.Wait()语句之后，可以看到程序正常运行，但是存在一个问题？
		两个go协程之间没有任何通信，相当于两个函数，
		但是在常规情况下，我们在创建协程之后，是希望协程之间实现“资源共享、通信计算”的。
	*/
	// 比如统计一共数过多少动物？
	/*
		可以想到的解决方案就是定义一个全局变量，每次调用count时，全局变量+1，但是这样会导致一个问题
		不同的go协程会被分配为不同的cpu，那么全局变量可能被不同的cpu同时操作
		一般编程语言操作方法：count调用全局变量时，给全局变量上锁，每次只能被一个CPU操作
		GO语言操作方法：（我们不通过共享内存来交流，我们要通过交流来共享内存）

	*/
}
