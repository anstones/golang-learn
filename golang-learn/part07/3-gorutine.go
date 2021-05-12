package main




import (
	"fmt"
	"time"
)

func gorutineTest(){
	i := 0
	for {
		i ++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(time.Second)
	}
}




func main(){
	go gorutineTest()
	i := 0
	for {
		i ++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i ==2{
			break   // 如果主协程退出了，其他任务将不执行
		}
	}
}

