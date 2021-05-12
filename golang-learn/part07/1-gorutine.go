package main

// 启动一个协程

import (
	"fmt"
	"time"
)

func hello()  {
	fmt.Println("Hello Goroutine! ")
}


func main() {
	go hello()
	fmt.Println("main Goroutine done")
	time.Sleep(time.Second)
}

