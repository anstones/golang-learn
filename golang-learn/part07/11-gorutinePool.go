package main

// goroutine池 可以有效控制goroutine数量，防止暴涨

import (
"fmt"
"math/rand"
)

type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

// 计算结果，发送到队列
func Calculation(jobChan chan *Job, resultChan chan *Result) {
	for job := range jobChan {
		// 随机数接过来
		r_num := job.RandNum
		// 随机数每一位相加
		// 定义返回值
		var sum int
		for r_num != 0 {
			tmp := r_num % 10
			sum += tmp
			r_num /= 10
		}
		// 想要的结果是Result
		r := &Result{
			job: job,
			sum: sum,
		}
		//运算结果扔到管道
		resultChan <- r
	}
}


// 从队列取出结果，打印
func Output(resultChan chan *Result) {
	for res := range resultChan{
		fmt.Printf("job id:%v randnum:%v result:%d\n", res.job.Id, res.job.RandNum, res.sum)
	}
}


// 创建工作池  参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go Calculation(jobChan, resultChan)
	}
}


func main() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go Output(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

