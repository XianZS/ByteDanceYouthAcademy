/*
map数据类型：

	(相当于python的dict类型)
	[1]语法：变量名:=make(map[key_type]value_type)
		key_type:表示字典key的数据类型
		value_type:表示字典value的数据类型
	[2]访问方式:
		变量名[key]:
			① 当key存在时:返回对应的value
			② 当key不存在时:根据type(value)返回一个固定的数值
				如type(value)=int,返回0
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("map数据类型")
	nameAge := make(map[string]int)
	nameAge["jom"] = 16
	nameAge["kom"] = 48
	nameAge["lom"] = 39
	fmt.Println(nameAge["aom"])
}
