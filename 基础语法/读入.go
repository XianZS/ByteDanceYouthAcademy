/*
[1]因为Scanln()遇到换行符或EOF的时候终止读取，所以在输入的时候只要按下回车键就会结束读取.
[2]Scanf可指定分隔符，其中上面的是%s %d中间的空格就是分隔符。例如下面指定:作为分隔符.
[3]Scanf可指定分隔符，其中上面的是%s %d中间的空格就是分隔符。例如下面指定:作为分隔符.
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println("读入")
	var number int
	var name string
	_, err := fmt.Scan(&number, &name)
	fmt.Println("number:", number)
	fmt.Println("name:", name)
	fmt.Println(err)
}
