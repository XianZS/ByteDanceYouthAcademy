package main

import (
	"fmt"
)

func Change(myArray [10]int) int {
	// 这种传参方式属于【子拷贝方式】
	for index, num := range myArray {
		fmt.Println(index, ":", num)
	}
	return 0
}

func Change1(myArray [10][10]int) int {
	for index, num := range myArray {
		fmt.Println(index, ":", num)
	}
	myArray[1][1] = 1000
	return 1
}

func main() {
	fmt.Println("动态数组-切片数组")
	var nums [10]int
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	for index, value := range nums {
		fmt.Println(index, "", value)
	}
	// %T的使用
	fmt.Printf("%T", nums)
	Change(nums)
	var news [10][10]int
	fmt.Println(news[1][1]) // 0
	Change1(news)
	fmt.Println(news[1][1]) // 0
	// 可以看到go语言函数传入参数属于深拷贝

}
