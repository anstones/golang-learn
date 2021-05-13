package main

import (
	"fmt"
	"strconv"
	"time"
)

// 用于判断管道是否存满

func writer(ch chan string)  {
	for i := 0; i<100; i++{
		select {
		case ch <- "hello" + strconv.Itoa(i):
			fmt.Println("write hello ", i)
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond*500)
	}
}

func main() {
	output := make(chan string, 10)

	go writer(output)

	for s := range output{
		fmt.Println("res: ", s)
		time.Sleep(time.Second)
	}


}