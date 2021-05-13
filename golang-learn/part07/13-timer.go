package main


// Ticker：时间到了，多次执行

import (
	"fmt"
	"time"
)


func Ticker(i int, ticker *time.Ticker)  {
	for  {
		i ++
		fmt.Println(<-ticker.C)
		if i == 5{
			ticker.Stop()
		}
	}
}


func main()  {
	ticker := time.NewTicker(1*time.Second)
	i := 0
	go Ticker(i, ticker)
	for  {

	}
}
