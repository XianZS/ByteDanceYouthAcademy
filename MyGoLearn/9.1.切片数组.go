package main

import "fmt"

func PrintfArray(news []int) int {
	news[1] = 100
	for _, num := range news {
		fmt.Println(num)
	}
	return 0
}

func main() {
	fmt.Println("切片数组")
	// 这就是所谓的(动态数组)切片数组
	nums := []int{1, 2, 3}
	fmt.Println(nums)                            // [1 2 3]
	fmt.Printf("%T\n", nums)                     // []int
	fmt.Println("nums[1]的初始数值为:", nums[1]) // 2
	PrintfArray(nums)
	fmt.Println("PrintfArray执行之后nums[1]的值为:", nums[1]) // 10
}
