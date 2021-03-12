package main

import "fmt"

//  通过匿名字段，可获得和继承类似的复用能力。


type Users struct {
	id int
	name string
}


type Mansger struct {
	Users
	title string
}



func (self *Users) ToString () string{
	return fmt.Sprintf("User: %p, %v", self, self)
}


func(self *Mansger) ToString() string{
	return fmt.Sprintf("Manager : %p, %v", self, self)
}

func main() {
	m := Mansger{Users{11, "sss"}, "M"}
	fmt.Println(m.ToString())
	fmt.Println(m.Users.ToString())
}
