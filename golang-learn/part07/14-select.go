package main

import (
	"fmt"
	"time"
)

// select可以同时监听一个或多个channel，直到其中一个channel ready

func selectTest1(ch chan string)  {
	time.Sleep(time.Second*2)
	ch <- "test1"
}

func selectTest2(ch chan string)  {
	time.Sleep(time.Second*5)
	ch <- "test2"
}

func main() {
	output1 := make(chan string)
	outpur2 := make(chan  string)

	go selectTest1(output1)
	go selectTest2(outpur2)

	select {
	case s1 := <- output1:
		fmt.Println(s1)
	case s2 := <- outpur2:
		fmt.Println(s2)
	}


}



