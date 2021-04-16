package main

import "fmt"

// 多个类型实现统一接口

type MoverFast interface {
	move()
}

type dogfast struct {
	name string
}


type Car struct {
	brand string
}

func (c Car)move()  {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func (d dogfast)move() {
	fmt.Printf("%s会跑\n", d.name)
}


func main() {
	var x MoverFast
	var a = dogfast{"旺财"}
	var b = Car{"大众"}

	x = a
	x.move()

	x = b
	x.move()


}
