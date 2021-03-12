package main

import "fmt"

type Person struct {
	name string
	age  int
}

// NewPerson 构造函数
func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p *Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

// SetAge 设置p的年龄
// 使用指针接收者
func (p *Person) SetAge(newage int) {
	p.age = newage
}

// SetAge2 设置p的年龄
//使用值接收者
func (p Person) SetAge2(newage int) {
	p.age = newage
}

func main() {
	// p01 := NewPerson("小李", 20)
	// p01.Dream()
	// fmt.Println(p01.age)
	// p01.SetAge(25)
	// fmt.Println(p01.age)

	p02 := NewPerson("小李", 20)
	p02.Dream()
	fmt.Println(p02.age)
	p02.SetAge2(25)
	fmt.Println(p02.age)
}
