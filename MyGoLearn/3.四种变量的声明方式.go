package main

/*
前三种方法可以进行全局变量声明，第四种方法不能进行全局变量声明。
*/

import (
	"fmt"
)

func main() {
	fmt.Println("四种变量的声明方式")
	// 1.声明一个变量
	var a int
	fmt.Println("a默认数值为:", a)
	// 2.声明一个变量，初始化一个数值
	var b int = 100
	fmt.Println("b初始化为:", b)
	// 3.在初始化的时候，可以省去数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c初始化为:", c)
	// 4.直接:=赋值
	// [:=]的缺陷，只能用在函数体内进行声明，不可以进行全局变量声明
	d := 100
	fmt.Println("d初始化为:", d)

	/*
		声明多个变量
	*/
	var x, y int = 100, 200
	fmt.Println("x:", x, ",y:", y)
	var (
		vv int = 100
		xx int = 200
	)
	fmt.Println(vv, xx)
}
