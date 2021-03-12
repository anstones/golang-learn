package main

import (
	"errors"
	"fmt"
)

// 标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象

var errDivByZero = errors.New("devision by zero")

func div(x,y int)(int, error) {
	if y == 0{
		return 0, errDivByZero
	}
	return x/y, nil
}

func main() {
	defer func() {
		fmt.Println(recover())
	}()

	z, err := div(10, 0)
	switch err {
	case nil:
		fmt.Println(z)
	case errDivByZero:
		panic(err)
	}
}
