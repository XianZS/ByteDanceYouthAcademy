package main

import (
	"fmt"
)

func main() {
	fmt.Println("这是switch语句")
	// 每个case默认break，不用像C/C++语言一样，在case之后写break
	var some = 1
	switch some {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("4")
	}
}
