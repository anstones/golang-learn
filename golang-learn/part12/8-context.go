package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 标准库context，它定义了Context类型，专门用来简化对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作


var wg sync.WaitGroup

func worker(ctx context.Context)  {
	go worker2(ctx)
	for {
		fmt.Println("worker")
		select {
		case <- ctx.Done(): // 等待通知
			break
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context)  {
	for {
		fmt.Println("worker2")
		time.Sleep(1*time.Second)
		select {
		case <- ctx.Done():
			break
		default:
		}
	}

}

func main() {
	ctx, channel := context.WithCancel(context.Background())
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second*3)
	channel()   // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}