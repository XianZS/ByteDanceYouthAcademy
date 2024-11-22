package main

import "fmt"

func main() {
	fmt.Println(" map 基础使用 - 2")
	nums := map[string]int{
		"jom": 2,
		"kom": 3,
		"lom": 4,
	}
	// 添加
	nums["aom"] = 5
	// 遍历
	fmt.Println(nums) // map[aom:5 jom:2 kom:3 lom:4]
	// 删除
	delete(nums, "jom")
	// 修改
	nums["kom"] = 99
	// 遍历
	for key, value := range nums {
		fmt.Print(key, " : ", value, " ")
	}
	// kom : 99 lom : 4 aom : 5
}
