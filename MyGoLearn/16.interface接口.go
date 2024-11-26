package main

import "fmt"

// interface{} 是万能类型
func myFunc(arg interface{}) {
	// 断言类型
	switch v := arg.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	case bool:
		fmt.Println("Boolean", v)
	case Books:
		fmt.Println("Book", v)
	default:
		fmt.Println("Don't know type")
	}
}

type Books struct {
	auth string
}

func main() {
	book1 := Books{auth: "auth1"}
	fmt.Println(book1.auth) // auth1
	myFunc(book1.auth)      // String auth1
}
