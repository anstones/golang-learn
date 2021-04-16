package main

//一个接口的方法，不一定需要由一个类型完全实现，接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}


type dryer struct {}

func (d dryer)dry()  {
	fmt.Println("甩一甩")
}

type haier struct {
	dryer
}

func (h haier)wash()  {
	fmt.Println("洗一洗")
}

func main() {
	var ww  WashingMachine
	var hh = haier{
		dryer{},
	}
	ww = hh
	ww.wash()
	ww.dry()
}