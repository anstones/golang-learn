package main

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

