package main

import (
	"fmt"
)

func add1(n int) {
	// 值类型传递
	n++
}

func add2(n *int) {
	// 指针类型传递
	*n++
}

func Pr(n int) {
	fmt.Println("n : ", n)
}

func Chance(nums []int) {
	// 应用类型传递
	nums[0] = -1
}

func main() {
	fmt.Println("指针")
	n := 2
	Pr(n)
	add1(n)
	Pr(n)
	add2(&n)
	Pr(n)
	nums := make([]int, 3)
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	fmt.Println(nums)
	Chance(nums)
	fmt.Println(nums)
}
