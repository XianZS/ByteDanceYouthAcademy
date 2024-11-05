package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

// 目标文件目录名称
var query = "test"

// 全局变量：承担计数器的作用
var matches int

// 当前正在工作的工人数
var workerCount = 0

// 最大工人数
var maxWorkerCount = 32

/*
	创建chan，通过chan交流来实现共享内存
*/
// [1]第一个chan：string，表示当前分配给工人的工作目录，也就是工作名
var searchRequest = make(chan string)

// [2]第二个chan：bool，表示当前的工作是否完成
var workerDone = make(chan bool)

// [3]第三个chan：bool，表示关于是否找到搜索结果
var foundMatch = make(chan bool)

func main() {
	start := time.Now()
	workerCount = 1
	go search("C:/", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
	/*
		【1】未优化结果:
			50 matches1
			19.3153321s
		【2】优化结果:
			50 matches
			1.4645552s
	*/
}
func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			// 存在新的路径(活)，指派工人
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundMatch <- true
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- path + name + "/"
				} else {
					go search(path+name+"/", false)
				}
			}
		}
		if master {
			workerDone <- true
		}
	}
}
