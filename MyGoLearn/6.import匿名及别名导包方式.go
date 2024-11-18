package main

import (
	// 如果导入某个不使用的包，则可以在其前加上“下划线 _ ”
	// 1、匿名方式:{语法: _ 包，在这种方式下，init()依旧会被调用，但是允许函数在main中不适用该包}
	_ "ByteDanceYouthAcademy/MyGoLearn/lib5/lib1"
	// 2、MyLib 别名方式:{语法: 别名 包}
	MyLib "ByteDanceYouthAcademy/MyGoLearn/lib5/lib2"
	// 3、直接导入包中的开放方法:{语法: . 包，这样就可以直接使用包中的开放语法，不建议轻易使用}
	. "ByteDanceYouthAcademy/MyGoLearn/lib5/lib1"
	"fmt"
)

func main() {
	fmt.Println("匿名别名导包方式")
	MyLib.Lib2()
	Lib1()
}
