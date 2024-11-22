package main

import "fmt"

func changeValue1(p int) {
	p = 10
}

func main1() {
	var a int = 1
	// 值传递
	changeValue1(a)
	fmt.Println("a = ", a)
}

func changeValue2(p *int) {
	// 【P *int】表示p存储的是地址，允许使用*p来改变p所指向的地址的内容
	*p = 10
}
func main2() {
	var a int = 1
	// 地址传递
	changeValue2(&a)
	fmt.Println("a = ", a)
}

func changeValue3(a *int, b *int) {
	number := *a
	*a = *b
	*b = number

}

func main3() {
	a := 10
	b := 20
	changeValue3(&a, &b)
	fmt.Println("a=", a, ",", "b=", b)
}

func main() {
	main1()
	main2()
	main3()
	var number int = 10
	var p *int
	p = &number
	fmt.Println(number)  // 10
	fmt.Println(&number) // 0xc00000a120
	fmt.Println(p)       // 0xc00000a120
	fmt.Println(*p)      // 10
}
