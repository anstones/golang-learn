package main

import "fmt"

// 空接口



// 空接口接收参数
func show(a interface{})  {
	fmt.Printf("type: %T, value: %v\n ", a, a)
}



// 空接口作为map的值, 使用空接口实现可以保存任意值的字典。
func interface_map(){
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "林白"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}


func main() {

	// 定义空接口
	var x interface{}
	x = "pprof.cn"
	fmt.Printf("type:%T value:%v\n", x, x)

	x = 100
	fmt.Printf("type:%T value:%v\n", x, x)

	x = true
	fmt.Printf("type:%T value:%v\n", x, x)

	show(x)

	interface_map()
}