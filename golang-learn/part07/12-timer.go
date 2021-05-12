package main

import (
	"fmt"
	"time"
)

// 定时器

func main() {
	timer1 := time.NewTimer(2*time.Second)
	fmt.Printf("t1:%v\n", time.Now())
	t2 := <- timer1.C
	fmt.Printf("t2:%v\n", t2)
}