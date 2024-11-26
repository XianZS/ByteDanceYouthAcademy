package main

import "fmt"

func main() {
	var a string
	// pair<static type :string,value :"jom">
	a = "jom"
	var b interface{}
	// pair<static type :interface{},value :jom>
	b = a
	fmt.Println(b)
	v, _ := b.(string)
	fmt.Println(v) // jom
}
