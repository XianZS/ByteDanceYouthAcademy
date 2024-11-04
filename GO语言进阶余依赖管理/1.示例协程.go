/*
协程：

	[1]语法：go func(传入参数A 传入参数A类型,传入参数B 传入参数B类型,……){
				函数体
			}(传入参数A 传入参数A类型,传入参数B 传入参数B类型,……)
*/
package main

import (
	"fmt"
	"time"
)

func Pr(i int) {
	fmt.Println("func Pr() : ", i)
}

func main() {
	fmt.Println("示范协程")
	for i := 0; i < 5; i++ {
		go func(j int) {
			Pr(j)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
