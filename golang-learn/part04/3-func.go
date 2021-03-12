package main

import "fmt"

// 普通函数与方法的区别


func valueTest(a int )int{
	return a +10
}


func pointerTest(a *int) int{
	return *a + 10
}

func structTestValue()  {
	a := 2
	fmt.Println("valueIntTest: ", valueTest(a))
	//函数的参数为值类型，则不能直接将指针作为参数传递
	//fmt.Println("valueIntTest:", valueIntTest(&a))

	b := 5
	fmt.Println("pointerIntTest: ", pointerTest(&b))
	//同样，当函数的参数为指针类型时，也不能直接将值类型作为参数传递
	//fmt.Println("pointerIntTest:", pointerIntTest(b))
}



type PersonID struct {
	id int
	name string
}

func (p PersonID)valueShowNmae()  {
	fmt.Println(p.name)
}

func(p *PersonID)pointShowName(){
	fmt.Println(p.name)
}


func structTestFunc(){

	//值类型调用方法
	personvalue := PersonID{100, "sssss"}
	personvalue.pointShowName()
	personvalue.valueShowNmae()


	//指针类型调用方法
	personpointer := &PersonID{20, "111111"}
	personpointer.valueShowNmae()
	personpointer.pointShowName()

	//与普通函数不同，接收者为指针类型和值类型的方法，指针类型和值类型的变量均可相互调用
}

func main() {
	structTestValue()
	structTestFunc()
}