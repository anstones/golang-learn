package main

import "fmt"

// 一个类型实现多个接口


type SayerNew interface {
	say()
}

type Mover interface {
	move()
}


type dog struct {
	name string
}

func(d dog) say(){
	fmt.Printf("%s会叫汪汪汪\n", d.name)
}


func(d dog)move(){
	fmt.Printf("%s会动\n", d.name)
}

func main()  {
	var x SayerNew
	var y Mover

	var wangcai = dog{"旺财"}

	x = wangcai
	y = wangcai
	x.say()
	y.move()
}

