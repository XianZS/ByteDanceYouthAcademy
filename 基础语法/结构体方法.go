/*
[1]已知结构体stu(name,id)
[2]结构体方法语法：

	① 修改原结构体
	func (s *stu) 方法名(传入参数 传入参数类型) 返回值类型{
		方法体
	}
	② 不修改原结构体
	func (s *stu) 方法名(传入参数 传入参数类型) 返回值类型{
		方法体
	}
*/
package main

import (
	"fmt"
)

type stu struct {
	name string
	num  int
}

func (s *stu) setNum(num int) bool {
	s.num = num
	return true
}
func main() {
	fmt.Println("结构体方法")
	stu1 := stu{name: "kom", num: 1}
	fmt.Println(stu1) // {kom 1}
	stu1.setNum(10)
	fmt.Println(stu1) // {kom 10}
}
