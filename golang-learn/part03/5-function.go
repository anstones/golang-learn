package main

import "fmt"

//defer 延迟调用

/*
defer特性：
    1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
defer用途：
    1. 关闭文件句柄
    2. 锁资源释放
    3. 数据库连接释放
*/


func test_defer_1(){
	var whatever [5]struct{}
	for i := range whatever{
		defer fmt.Println(i)
	}
}

// defer 碰上闭包
func test_defer_2(){
	var whatever [5]struct{}
	for i := range whatever{
		fmt.Printf("before defer %d\n",i)
		defer func() {fmt.Println(i)}()
	}

}


func main() {
	//test_defer_1()
	test_defer_2()
}





