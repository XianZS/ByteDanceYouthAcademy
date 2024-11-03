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
	[3]删除操作:
		delete(变量名,key)
	[4]返回值:
		假设nums为map[string]int{"1":1,"2":2,"3":3}
		map会返回两个值:
			* 第一个返回值是你要查找的数据(存在返回数据，不存在返回默认)
			* 第二个返回值是bool类型，如果你查找的数值存在，则为true，反之为false
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
	fmt.Println(nameAge)
	// map删除操作
	delete(nameAge, "lom")
	fmt.Println(nameAge)
}
