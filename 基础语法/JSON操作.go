/*
在进行JSON操作时，需要保证结构体的公开字段(属性变量名)首字母大写
[1]JSON序列化操作语法
res1,err1 := json.Marshal(序列化对象)
fmt.Println(res1,err1)
[2]JSON反序列化操作语法
var res2 序列化对象类型
err2 = json.Unmarshal(JSON对象res1,反序列化存储对象res2)
fmt.Println(res2,err2)
*/
package main

import (
	"encoding/json"
	"fmt"
)

type MyStruct struct {
	Name string
	Id   int
}

func main() {
	fmt.Println("JSON操作")
	myStruct := MyStruct{"jom", 21}
	some, err := json.Marshal(myStruct)
	fmt.Println(string(some), err)
	var myStruct2 MyStruct
	err = json.Unmarshal(some, &myStruct2)
	fmt.Println(myStruct2, err)
}
