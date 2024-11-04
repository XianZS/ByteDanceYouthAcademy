/*
字符串工具函数：
[1]字符是否存在：strings.Contains(字符串，被查找字符)
[2]字符的下标：strings.Index(字符串，被查找字符)
[3]字符串计数：strings.Count(字符串，被查找字符)
[4]字符串连接函数：strings.Join([]string{被链接子集}，连接符号)
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("字符串工具函数")
	a := "hello"
	fmt.Println(strings.Contains(a, "l"))
	fmt.Println(strings.Index(a, "l"))
	fmt.Println(strings.Count(a, "l"))
	fmt.Println(strings.Join([]string{"a", "b", "c"}, "-"))

}
