package main

import "fmt"

// 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。
func pointer1() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

// 空指针的判断
func pointer2() {
	var p *string
	fmt.Println(p)
	fmt.Printf("P的值是%v", p)
	if p != nil {
		fmt.Println("非空")
	} else {
		fmt.Println("空值")
	}
}

//程序定义一个int变量num的地址并打印
// 将num的地址赋给指针ptr，并通过ptr去修改num的值
func pointer3() {
	var a int
	fmt.Println(&a)
	var p *int
	p = &a
	*p = 20
	fmt.Println(a)
}

func main() {
	// pointer1()
	// pointer2()
	pointer3()
}
