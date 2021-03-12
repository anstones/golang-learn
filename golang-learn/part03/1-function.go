package main

import "fmt"

// 函数参数的引用传递

func swap(x, y *int)  {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}


// 不定长参数传参
func test(s string, n ...int) string {
	var x int
	for _, i := range n{
		x += i
	}
	return  fmt.Sprintf(s, x)
}


func main() {
	var a,b int = 1,2
	swap(&a, &b)
	fmt.Println(a, b)

	fmt.Println(test("1111111"))
	fmt.Println(test("sum: %d", 1,2,3)) //个或者多个

	var slice = []int{1,2,3}
	fmt.Println(test("sum:%d", slice...)) // 不定长参数传递slice对象时，必须展开
}