package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("进程信息")
	fmt.Println(os.Args) // 进程在执行时的命令行参数
	// [C:\Users\Administrator\AppData\Local\JetBrains\GoLand2024.2\tmp\GoLand\___3go_build__go.exe]
	fmt.Println(os.Getenv("GOPATH")) // 获取GO的PATH
	// D:\program\GoLand\GOPATH
	buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts/").CombinedOutput() // 快速启动子进程，并获取相关信息
	//[] exec: "grep": executable file not found in %PATH%
	fmt.Println(buf, err)
}
