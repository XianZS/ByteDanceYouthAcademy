// using utf8
package main

import (
	"fmt"
)

// 通过 const 定义枚举类型
const (
	name    = "Jom"
	number  = 999
	address = "北京"
)

func main() {
	fmt.Println("常量注意事项")
	// 定义一个常量，只读属性
	const number int = 10

}
