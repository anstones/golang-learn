package main

import "fmt"

// 指针类型匿名字段

type Person struct {
	name string
	sex string
	age int
}


type Student struct {
	*Person
	id int
	addr string
}

func main()  {
	s1 := Student{&Person{"s", "女", 25}, 1, "2222"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)
}
