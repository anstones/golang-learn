package main

import (
	"fmt"
)

// 如果多个channel同时ready，则随机选择一个执行

func main() {
	int_chan := make(chan int)
	string_chan := make(chan string)

	go func() {
		string_chan <- "hello"
	}()
	
	go func() {
		int_chan <- 1
	}()

	select {
	case value := <- string_chan:
		fmt.Println("string:", value)
	case value := <- int_chan:
		fmt.Println("int:", value)
	}
}