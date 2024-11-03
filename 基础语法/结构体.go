/*
[1]结构体语法:

	type 变量名 struct{变量名 数据类型}

示范案例:

	type stu struct{
		name string
		id int
		chinese int
	}

[2]创建结构体
*/
package main

import "fmt"

type user struct {
	name string
	id   int
}

func chance1(a user) bool {
	a.name = "测试"
	return true
}

func chance2(a *user) bool {
	a.name = "测试"
	return true
}
func main() {
	stu1 := user{name: "kom", id: 1}
	fmt.Println("stu1 : ", stu1) // stu1 :  {kom 1}
	// 值类型传递
	chance1(stu1)
	fmt.Println("stu1 : ", stu1) // stu1 :  {kom 1}
	// 引用类型传递
	chance2(&stu1)
	fmt.Println("stu1 : ", stu1) // stu1 :  {测试 1}
}
