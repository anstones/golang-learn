package main

// 接口与接口间可以通过嵌套创造出新的接口。

import "fmt"

type Sayer interface {
	say()
}


type Mover interface {
	move()
}


type animal interface {
	Sayer
	Mover
}


type cat struct {
	name string
}

func (c cat)say(){
	fmt.Println("喵喵")
}

func (c cat) move(){
	fmt.Println("跳")
}

func main() {
	var x animal
	x = cat{"皮球"}
	x.move()
	x.say()


}