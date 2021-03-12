package main

import (
	"fmt"
	"strings"
)

/*
   1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
   2. 切片的长度可以改变，因此，切片是一个可变的数组。
   3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
   4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
   5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
   6. 如果 slice == nil，那么 len、cap 结果都等于 0。
*/

func slice1() {
	var s1 []int
	if s1 == nil {
		fmt.Println("空")
	} else {
		fmt.Println("不是空")
	}

	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2, 3, 4}
	fmt.Println(s5)

	arr := [5]int{1, 2, 3, 4, 5}
	s6 := arr[1:4]
	fmt.Println(s6)
}

// 读写操作实际目标是底层数组，只需注意索引号的差别。
func slice2() {
	data := []int{0, 1, 2, 3, 4, 5}

	s := data[2:4]

	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)
}

// 指针直接访问底层数组，退化成普通数组操作。
func slice3() {
	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100

	fmt.Println(s)
}

//用append内置函数操作切片（切片追加）
func slice4() {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s1, s2)
}

// 函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
func slice5() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice data ", data)

	s1 := data[8:]
	s2 := data[:5]

	fmt.Printf("slice s1: %v\n", s1)
	fmt.Printf("slice s2: %v\n", s2)

	copy(s1, s2)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", data)
}

// 切片便利

func slice6() {
	data := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	fmt.Printf("%v\n", slice)
	for index, value := range data {
		fmt.Printf("inde : %v , value : %v\n", index, value)
	}
}

// string本身是不可变的，因此要改变string中字符。需要如下操作： 英文字符串：
func slice7() {
	str := "hello world"
	s := []byte(str) //中文字符需要用[]rune(str)

	s[6] = 'G'
	s = s[:8]
	s = append(s, '!') // 双引号是用来表示字符串string，其实质是一个byte类型的数组，单引号表示rune类型。
	str = string(s)
	fmt.Println(str)
}

// 数组or切片转字符串：
func slice8() {
	array_or_slice := []string{"123", "hello go"}
	res := strings.Replace(strings.Trim(fmt.Sprint(array_or_slice), "[]"), " ", ",", -1)
	fmt.Println(res)
}

func main() {
	// slice1()
	// slice2()
	// slice3()
	// slice4()
	// slice5()
	// slice6()
	// slice7()
	slice8()
}
