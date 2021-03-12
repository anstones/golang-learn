package main

import "fmt"

// select可以监听channel的数据流动
// 与switch语句可以选择任何使用相等比较的条件相比，select由比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作

func main() {
	var c1, c2, c3 chan int

	var i1, i2 int

	select {
	case i1 = <- c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok:= (<-c3):
		if ok{
			fmt.Printf("received ", i3, " from c3\n")
		}else{
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}

}