package main

import "fmt"

// swith 类型判断

func switch_test_1(){
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型 :%T\r\n", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
	
}

func switch_test_2()  {
	var j = 1
	switch j {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

}

func switch_test_3()  {
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
		/*
		Go的switch非常灵活，表达式不必是常量或整数，执行的过程从上至下，直到找到匹配项；
		而如果switch没有表达式，它会匹配true。
		Go里面switch默认相当于每个case最后带有break，
		匹配成功后不会自动向下执行其他case，而是跳出整个switch,
		但是可以使用fallthrough强制执行后面的case代码。
		*/
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
}

func switch_test_4()  {
	var n =0
	switch{
	case n > 0 && n<10:
		fmt.Println("i > 0 and i < 10")
	case n > 10 && n < 20:
		fmt.Println("i > 10 and i < 20")
	default:
		fmt.Println("def")
	}
}

func main() {
	switch_test_2()
	switch_test_3()
}
