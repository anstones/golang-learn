package main

import (
	"fmt"
	"strconv"
)

// 类型装换

func AtoiDemo()  {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("can't convert to int")
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1) //type:int value:100
	}
}

func ItoaDemo()  {
	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Printf("type:%T value:%#v\n", s2, s2) //type:string value:"200"
}

func ParseDemo()  {
	// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
	//b, err := strconv.ParseBool("true")
	//f, err := strconv.ParseFloat("3.1415", 64)
	//i, err := strconv.ParseInt("-2", 10, 64)
	//u, err := strconv.ParseUint("2", 10, 64)
}

func FormatDemo()  {
	// Format系列函数实现了将给定类型数据格式化为string类型数据的功能。
	//s1 := strconv.FormatBool(true)
	//s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	//s3 := strconv.FormatInt(-2, 16)
	//s4 := strconv.FormatUint(2, 16)
}





func main() {
	//AtoiDemo()
	ItoaDemo()
}

