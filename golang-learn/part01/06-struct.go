package main

import "fmt"

type person struct {
	name string
	city string
	age  int
}

type student struct {
	name string
	age  int
}

func struct1() {
	// 类型定义
	type NewInt int

	// 类型别名
	type MyInt = int

	var a NewInt
	var b MyInt

	fmt.Printf("type of a : %T\n", a)
	fmt.Printf("type of b : %T\n", b)
}

func struct2() {
	var p1 person
	var per = new(person)
	fmt.Printf("%T\n", p1)

	fmt.Printf("%T\n", per)
	fmt.Printf("%+v\n", per)
	fmt.Printf("%#v\n", per)

	// 使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
	p2 := &person{}
	fmt.Printf("%T\n", p2)

	// 3.name = "小明"其实在底层是(*p3).name="小明"，这是Go语言帮我们实现的语法糖。
	p2.name = "小明"
	fmt.Printf("p2= %#v\n", p2)

}

func struct3() {
	// 结构体初始化
	var p3 person
	fmt.Printf("p3=%#v\n", p3)

	//使用键值对初始化
	p4 := person{
		name: "xiaoming",
		city: "西安",
		age:  10,
	}
	fmt.Printf("p4=%#v\n", p4)

	//使用值的列表初始化
	p5 := person{
		"小李",
		"北京",
		20,
	}
	fmt.Printf("p5=%#v\n", p5)
}

func struct4() {
	m := make(map[string]student)

	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = stu
	}

	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

//构造函数
func struct5(name, city string, age int) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func main() {
	// struct1()
	//struct2()
	//struct3()
	//struct4()
	p6 := struct5("xiaoming", "西安", 20)
	fmt.Printf("person : %#v\n", p6)

}
