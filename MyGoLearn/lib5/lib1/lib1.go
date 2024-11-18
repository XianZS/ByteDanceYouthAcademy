package lib1

import "fmt"

/*
	- 函数名首字母大写，表示当前函数对外开放，public
	- 函数名首字母小写，表示当前函数不对外开放，private
*/

func Lib1() {
	fmt.Println("Lib1()")
}

/*
- 对于每一包来讲，init会在编译时最先执行
- init()的执行优先级大于main()的执行优先级
*/
func init() {
	fmt.Println("lib1 init")
}
