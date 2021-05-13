package main

import (
	"fmt"
	"sync"
	"time"
)

// 场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，这种场景下使用读写锁是更好的一种选择


var (
	x int
	wg sync.WaitGroup
	lock sync.Mutex
	rwlock sync.RWMutex

)


func write(){
	rwlock.Lock()  // 加写锁
	x += 1
	time.Sleep(10*time.Millisecond)
	rwlock.Unlock()  // 解写锁
	wg.Done()
}

func rewd()  {
	rwlock.RLock()   // 加读锁
	time.Sleep(time.Millisecond)
	rwlock.RUnlock()  // 解读锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0;i<10;i++{
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go rewd()
	}

	wg.Wait()

	end := time.Now()

	fmt.Println(end.Sub(start))
}