package main

import "fmt"

//Go实现类似 try exception 的异常处理

func Try(fun func(), handler func(interface{}))  {
	defer func(){
		if err := recover(); err != nil{
			handler(err)
		}
	}()
	fun()
}

func main() {
	Try(func(){
		//panic("test panic")
		var i,j int = 4,0
		x := i/j
		fmt.Println(x)
	}, func(err interface{}) {
		fmt.Println(err)
	})
}

