package main

import "fmt"

type Book struct {
	title  string
	author string
}

func pr1(book Book) {
	book.title = "pr1"
}

func pr2(book *Book) {
	// 如果这个变量是指针的话，那么它存储的是地址
	// 如果变量是变量的话，那么它存储的是元素本身
	book.title = "pr2"
}

func main() {
	fmt.Println("Go语言结构体")
	var book1 Book
	book1.title = "Timer"
	book1.author = "kom"
	fmt.Println(book1) // {Timer kom}
	pr1(book1)
	fmt.Println(book1) // {Timer kom}
	pr2(&book1)
	fmt.Println(book1) // {pr2 kom}
}
