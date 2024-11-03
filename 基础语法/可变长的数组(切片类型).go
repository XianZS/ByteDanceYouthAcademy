package main

import (
	"fmt"
)

func main() {
	fmt.Println("切片")
	/*
		【1】创建可变长数组
		语法：数组名:=make([]数组类型,数组长度)
	*/
	nums := make([]string, 3)
	nums[0] = "切片第一个元素"
	nums[1] = "切片第二个元素"
	nums[2] = "切片第三个元素"
	fmt.Println("nums:", nums, "len(nums):", len(nums))
	// nums: [切片第一个元素 切片第二个元素 切片第三个元素] len(nums): 3
	/*
		【2】添加元素
		语法：A数组名=append(B数组名,添加元素)
		如果想要在原数组的基础上添加元素，那么[A数组名==B数组名]
		如果想要重新建立一个新的数组，那么[A数组名:=append(B数组名,添加元素)]
		旧数组的改变，并不会影响新数组，相当于开辟一个新的堆空间
	*/
	some := append(nums, "切片第四个元素")
	nums[2] = "测试"
	fmt.Println(nums) // [切片第一个元素 切片第二个元素 测试]
	fmt.Println(some) // [切片第一个元素 切片第二个元素 切片第三个元素 切片第四个元素]
	/*
		【3】copy操作
		语法：copy(新数组，老数组)
		可以看到，老数组的改变并不会影响新数组
	*/
	news := make([]string, len(some))
	copy(news, some)
	some[2] = "这是some数组第二个元素"
	fmt.Println(news) // [切片第一个元素 切片第二个元素 切片第三个元素 切片第四个元素]
	fmt.Println(some) // [切片第一个元素 切片第二个元素 这是some数组第二个元素 切片第四个元素]
	/*
		【4】切片操作
		语法：和python类似，左闭右开
	*/
	fmt.Println(some[1:2]) // [切片第二个元素]
}
