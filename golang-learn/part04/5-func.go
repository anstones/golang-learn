package main

import (
	"errors"
	"fmt"
)

// 系统抛出异常


func test01(){

	a := [5]int{1,2,3,4,5}
	a [1] = 123

	fmt.Println(a)

	index := 10
	a[index] = 10
	fmt.Println(a)
}

func getCircleArea(radius float32)(arez float32)  {
	if radius < 0 {
		panic("半径不能为负")
	}
	return 3.14 *radius*radius
}


func getCircleArea1(radius float32)( area float32, err error)  {
	if radius < 0 {
		// 构建异常对象
		err = errors.New("半径不能为负")
		return
	}
	area = 3.14 *radius*radius
	return
}


func test02()  {
	getCircleArea(-5)
}


func test03()  {
	i , err := getCircleArea1(-5)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(i)
	}
}




func main() {
	//test01()
	//test02()
	test03()
}