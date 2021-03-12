package main

/*
   1. 数组：是同一种数据类型的固定长度的序列。
   2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
   3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
   4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
   for i := 0; i < len(a); i++ {
   }
   for index, v := range a {
   }
   5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
   6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
   7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
   8.指针数组 [n]*T，数组指针 *[n]T。
*/

import (
	"fmt"
	"math/rand"
)

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello world", 4: "tom"}

// 一维数组
func array1() {
	a := [3]int{1, 2}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	d := []struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}

	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)
}

// 多维数组
var arr_0 [5][3]int
var arr_1 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

func array2() {
	a1 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	b2 := [...][3]int{{1, 2, 3}, {4, 5, 6}} // 第2 纬度不能用 "..."。

	fmt.Println(arr_0, arr_1)
	fmt.Println(a1, b2)
}

// 多维数组遍历
func array3() {
	f := [...][3]int{{1, 2, 3}, {7, 8, 9}}

	for k1, v1 := range f {
		for k2, v2 := range v1 {
			fmt.Printf("(%d, %d)= %d", k1, k2, v2)
		}
		fmt.Println()
	}
}

// 数组拷贝和传值

func printarr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func array4() {
	arr1_0 := [5]int{}
	printarr(&arr1_0)
	fmt.Println(arr1_0)
	arr2_0 := [...]int{2, 4, 6, 8, 10}
	printarr(&arr2_0)
	fmt.Println(arr2_0)

}

// 求数组所有元素之和
func sumarr(a [10]int) int {
	var sum int = 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func array5() {
	// rand.Seed(1)
	// rand.Seed(time.Now().Unix())

	var b [10]int
	for k, _ := range b {
		b[k] = rand.Intn(1000)
	}
	sum := sumarr(b)
	fmt.Printf("sum=%d\n", sum)
}

// 找出数组中和为给定值的两个元素的下标
func mytest(a [5]int, sum int) {
	for i := 0; i < len(a); i++ {
		other := sum - a[i]
		for j := 0; j < len(a); j++ {
			if a[j] == other {
				fmt.Printf("(%d, %d)", i, j)
			}
		}
	}
}

func array6() {
	arr_a := [5]int{1, 3, 5, 8, 7}
	mytest(arr_a, 8)
	fmt.Println()
}

func main() {
	// array1()
	// array2()
	// array3()
	// array4()
	// array5()
	array6()
}
