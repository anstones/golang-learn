package main

import (
	"fmt"
	"time"
)

// 定时器

func main() {
	//timer1 := time.NewTimer(2*time.Second)
	//fmt.Printf("t1:%v\n", time.Now())
	//t2 := <- timer1.C
	//fmt.Printf("t2:%v\n", t2)


	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<- timer2.C
	//	fmt.Println("时间到")
	//}

	// timer实现延时的功能
	//time3 := time.NewTimer(2*time.Second)
	//<-time3.C
	//fmt.Println("2秒到")
	//
	//<-time.After(2*time.Second)
	//fmt.Println("2秒到")

	//停止计时器
	//timer4 := time.NewTimer(2*time.Second)
	//go func() {
	//	<- timer4.C
	//	fmt.Println("定时器执行了")
	//}()
	//
	//b:= timer4.Stop()
	//if b{
	//	fmt.Println("timer4已经关闭")
	//}

	// 重置定时器
	timer5 := time.NewTimer(3*time.Second)
	timer5.Reset(1*time.Second)
	fmt.Println(time.Now())
	fmt.Println(<-timer5.C)
	for {

	}

}