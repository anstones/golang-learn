package main

import "fmt"

const pi = 3.1545

const (
	n1 = iota
	n2
	n3
	_
	n4
)

const (
	x1 = iota
	x2 = 100
	x3 = iota
	x4
)

func main() {
	fmt.Println("hello world")
	fmt.Println("unit0 is ", uint(0))

	fmt.Println("n4 ", n4)

	fmt.Println("x3 ", x3, "x4 ", x4)

}
