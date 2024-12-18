package main

import "fmt"

func main() {
	fmt.Println("动态数组的基础操作")

	/*
		【1】开辟动态数组
		语法：make([]type, 动态数组长度, 合法容量)
	*/
	var nums = make([]int, 2, 4)
	fmt.Println(nums, len(nums), cap(nums)) // [0 0] 2 4
	/*
		【2】给切片数组添加元素
		当数组目前存储的元素个数 === 数组的合法长度时，再次向数组添加元素，会发生什么？
		答：会再次开辟一个cap长度，如cap(nums)=4，那么再次开辟cap之后，cap(nums)=8
	*/
	nums = append(nums, 3)
	fmt.Println(nums, len(nums), cap(nums)) // [0 0 3] 3 4
	nums = append(nums, 4)
	fmt.Println(nums, len(nums), cap(nums)) // [0 0 3 4] 4 4
	nums = append(nums, 5)
	fmt.Println(nums, len(nums), cap(nums)) // [0 0 3 4 5] 5 8
	/*
		【3】len和cap的区别？
		len表示数组左右指针之间的距离，cap表示数组底层的长度
	*/

	/*
		【4】切片的截取
		操作规则：左闭右开
	*/
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)      // [1 2 3 4 5 6]
	fmt.Println(s[0:2]) // [1 2]

	/*
		【5】切片之后的切片数组和原切片数组有什么关系？
		切片之后的数组，本质上依旧指向原切片数组，更改两者任意其一，另外一个切片数组都会受到影响
	*/
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	newXs := xs[:4]
	fmt.Println(xs, newXs) // [1 2 3 4 5 6 7 8 9] [2 3 4]
	newXs[1] = 100
	fmt.Println(xs, newXs) // [1 2 100 4 5 6 7 8 9] [2 100 4]
	// 如何避免这种情况的发生？
	newXs_ := make([]int, 4, 10)
	copy(newXs_, xs[:4])
	fmt.Println(newXs_) // [1 100 3 4]
}
