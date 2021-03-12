package main

import "fmt"

type Address struct {
	Province   string
	City       string
	CreateTime string //指定Address结构体中的CreateTime
}

type Email struct {
	Account    string
	CreateTime string //指定Email结构体中的CreateTime
}

type User struct {
	Name    string
	Gender  string
	Address Address // 结构体嵌套
	Email
}

//通过嵌套匿名结构体实现继承
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	Feet int
	*Animal
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	user1 := User{
		Name:   "pprof",
		Gender: "女",
		Address: Address{
			Province: "黑龙江",
			City:     "哈尔滨",
		},
	}
	fmt.Printf("user1=%+v\n", user1)

	user2 := User{
		Name:   "xiaoming",
		Gender: "女",
		Address: Address{
			Province:   "黑龙江",
			City:       "哈尔滨",
			CreateTime: "12",
		},
		Email: Email{
			Account:    "143@163.com",
			CreateTime: "21",
		},
	}
	fmt.Println(user2.Address.CreateTime, user2.Email.CreateTime)

	d := Dog{
		Feet: 4,
		Animal: &Animal{
			name: "乐乐",
		},
	}
	d.move()
	d.wang()
}
