package main

//如何优雅的从通道循环取值


import "fmt"

func main()  {
	c1 := make(chan int)
	c2 := make(chan int)

	// 开启goroutine将0~100的数发送到c1中
	go func() {
		for i := 0; i<100; i++{
			c1  <- i
		}
		close(c1)
	}()

	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			if data, ok := <-c1;ok{
				c2 <- data
			}else {
				break
			}
		}
		close(c2)
	}()

	// 在主goroutine中从c2中接收值打印
	for i := range  c2{ // 通道关闭后会退出for range循环
		fmt.Println(i)
	}

}
