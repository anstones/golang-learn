package main

import "fmt"

// go 闭包

func add_(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

func main()  {
	temp1 := add_(10)
	// 此时tmp1是函数
	fmt.Printf("%T\n", temp1)
	fmt.Println(temp1(1), temp1(2))
}
