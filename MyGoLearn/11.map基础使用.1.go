package main

import "fmt"

func main() {
	fmt.Println(" map 基础使用 - 1")
	/*
		【1】map基础使用：
		map变量名 := make(map(key_type)value)type)
	*/
	myMap1 := make(map[string]string)
	fmt.Println(myMap1, len(myMap1))
	myMap1["name"] = "jom"
	myMap1["age"] = "19"
	myMap1["address"] = "172.0.0.1"
	fmt.Println(myMap1) // map[address:172.0.0.1 age:19 name:jom]
	myMap2 := make(map[int]map[string]string)
	myMap2[1] = map[string]string{"name": "jom"}
	myMap2[2] = map[string]string{"age": "19"}
	fmt.Println(myMap2) // map[1:map[name:jom] 2:map[age:19]]
}
