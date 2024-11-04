package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("时间处理")
	nowTime := time.Now()
	fmt.Println(nowTime)
	// 2024-11-04 09:10:57.5864502 +0800 CST m=+0.000000001
	// 获取当前时间点的一些信息
	fmt.Println(nowTime.Year())
	// 2024
	fmt.Println(nowTime.Month())
	// November
	fmt.Println(nowTime.Day())
	// 4
	fmt.Println(nowTime.Hour())
	// 9
	fmt.Println(nowTime.Minute())
	// 10
	fmt.Println(nowTime.Second())
	// 57
	fmt.Println(nowTime.Nanosecond())
	// 586450200
	fmt.Println(nowTime.Location())
	// Local
	fmt.Println(nowTime.YearDay())
	// 309
}
