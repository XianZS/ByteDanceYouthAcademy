/**
** 派生类对基类的虚方法进行重写
**/

package main

import "fmt"

type Animal interface {
	Sleep() string
	GetColor() string
	GetType() string
}

// 定义一个具体的类

type Cat struct {
	Color string
}

func (cat *Cat) Sleep() string {
	return "成功睡觉了"
}

func (cat *Cat) GetColor() string {
	return cat.Color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

func main() {
	// 那么在继承中要构成多态还有两个条件：
	// 1、必须通过基类的指针或者引用调用虚函数
	// 2、被调用的函数必须是虚函数，且派生类必须对基类的虚函数进行重写
	var cat1 Animal
	cat1 = &Cat{Color: "red"}
	fmt.Println(cat1)            // {red}
	fmt.Println(cat1.Sleep())    // 成功睡觉了
	fmt.Println(cat1.GetColor()) // red
	fmt.Println(cat1.GetType())  // Cat
}
