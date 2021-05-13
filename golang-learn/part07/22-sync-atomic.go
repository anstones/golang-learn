package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 代码中的加锁操作因为涉及内核态的上下文切换会比较耗时、代价比较高。
//针对基本数据类型我们还可以使用原子操作来保证并发安全，因为原子操作是Go语言提供的方法它在用户态就可以完成


var x int64
var lock sync.Mutex
var wg sync.WaitGroup

// 普通版加函数
func Add()  {
	x++
	wg.Done()
}

// 互斥锁版加函数
func mutexAdd(){
	lock.Lock()
	x++
	lock.Unlock()
	wg.Done()
}


// 原子操作版加函数
func atomicAdd()  {
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	for i:=0; i<5000;i++{
		wg.Add(1)
		//go Add()
		//go mutexAdd()
		go atomicAdd()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}