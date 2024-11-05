package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func use8() {
	some := ioutil.Discard
	fmt.Println(some)
	time.Sleep(time.Second)
}

var query1 = "test"
var matches1 int

func main() {
	start := time.Now()
	search7("C:/")
	fmt.Println(matches1, "matches1")
	fmt.Println(time.Since(start))
	/*
		50 matches1
		19.3153321s
	*/
}

func search7(path string) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query1 {
				matches1++
			}
			if file.IsDir() {
				search7(path + "/" + name)
			}
		}
	}
}
