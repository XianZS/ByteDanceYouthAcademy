package main

import "fmt"

type People struct {
	name    string
	age     int
	Address string
}

func (p *People) SayHello() {
	fmt.Printf("People Hello !")
}

type SuperPeople struct {
	People // 表示 SuperPeople 类继承了 People 类的方法
	grade  string
}

func (p *SuperPeople) SayHello() {
	fmt.Println("SuperPeople Hello")
}

func main() {
	p1 := People{name: "jom", age: 21, Address: "127.0.0.1"}
	fmt.Println(p1)                                         // {jom 21 127.0.0.1}
	p1.SayHello()                                           // Hello, jom!
	p2 := SuperPeople{People{"kom", 46, "127.0.0.2"}, "ds"} // People Hello !127.0.0.2

	fmt.Println(p2.Address)
}
