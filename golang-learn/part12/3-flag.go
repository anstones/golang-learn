package main

import (
	fmt1 "fmt"
	"os"
)

// 获取命令行参数

func main() {
	if len(os.Args) >0 {
		for index, arg := range os.Args{
			fmt1.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
