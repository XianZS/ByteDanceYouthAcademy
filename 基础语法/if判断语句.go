package main

import (
	"fmt"
)

func main() {
	fmt.Println("判断语句")
	var n = 1
	if n&1 == 1 {
		fmt.Println("这是一个奇数")
	} else {
		fmt.Print("这是一个偶数")
	}
}
