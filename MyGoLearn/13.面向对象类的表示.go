/*
go语言的类：

	通过结构体绑定方法
*/
package main

import "fmt"

type Hero struct {
	name string
	age  int
}

func (s Hero) GetName() {
	fmt.Println("Name = ", s.name)
}
func main() {
	fmt.Println("面向对象类的表示")
	// 创建一个对象
	hero := Hero{name: "jom", age: 10}
	fmt.Println(hero)
}
