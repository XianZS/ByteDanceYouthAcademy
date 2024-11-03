package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2))
	number := add(1, 2)
	fmt.Println(number)
	nums := map[string]int{"1": 1, "2": 2, "3": 3}
	fmt.Println(nums)
	var x int
	var y bool
	x, y = nums["4"]
	fmt.Println(x, y)
}
