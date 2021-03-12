package main

import "fmt"

type User struct {
	Name string
	Email string
}

func (u User)Notify()  {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

func (u *User)Notify1()  {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}


func main()  {
	u1 := User{"stone", "stone@163.com"}
	u1.Notify()

	u2 := User{"golang", "go@go.com"}
	//u3 := &u2
	//u3.Notify1()
	u2.Notify1()
}



