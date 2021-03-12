package main

import (
	"fmt"
	"time"
)

/*
select是Go中的一个控制结构，类似于switch语句，用于处理异步IO操作。select会监听case语句中channel的读写操作，当case中channel读写操作为非阻塞状态（即能读写）时，将会触发相应的动作。
select中的case语句必须是一个channel操作
select中的default子句总是可运行的。
如果有多个case都可以运行，select会随机公平地选出一个执行，其他不会执行。
如果没有可运行的case语句，且有default语句，那么就会执行default的动作。
如果没有可运行的case语句，且没有default语句，select将阻塞，直到某个case通信可以运行
*/

// 典型用法1 ： 判断超时
var reschan = make(chan int)

func doData(data int)  {
	// ...
}

//比如在下面的场景中，使用全局resChan来接受response，如果时间超过3S,resChan中还没有数据返回，则第二条case将执行
func select_timeout()  {
	select {
	case data := <- reschan:
		doData(data)
	case <- time.After(time.Second*3):
		fmt.Println("request time out")
	}
}


// 典型用法2 : 退出

var shouldQuit = make(chan struct{})

func select_out()  {
	{
		// loop
	}
	// out of the loop
	select {
	case <- shouldQuit:
		//cleanUp()
		return
	default:
	}
}

//再另外一个协程中，如果运行遇到非法操作或不可处理的错误，就向shouldQuit发送数据通知程序停止运行
//close(shouldQuit)


// 典型用法2 : 判断channel是否阻塞

//在某些情况下是存在不希望channel缓存满了的需求的，可以用如下方法判断
func main()  {
	ch := make(chan int, 5)
	data := 0
	select {
	case ch <- data:
	default:
		//做相应操作，比如丢弃data。视需求而定
	}
}