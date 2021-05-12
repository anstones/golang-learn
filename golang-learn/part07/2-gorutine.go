package main

import (
	"fmt"
	"sync"
)

// 启动多个goroutine

var ws sync.WaitGroup


func hello1(i int)  {
	defer ws.Done()    // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)

}

func main() {
	for i:=0; i<10; i++{
		ws.Add(1) // 启动一个goroutine就登记+1
		go hello1(i)
	}
	ws.Wait() // 等待所有登记的goroutine都结束
}