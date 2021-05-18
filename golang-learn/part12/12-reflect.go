package main

import (
	"fmt"
	"reflect"
)

// 反射是指在程序运行期对程序本身进行访问和修改的能力

//反射获取interface类型信息
func reflectType(a interface{})  {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)

	k := t.Kind()
	fmt.Println(k)

	switch k {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}


// 反射获取interface值信息
func reflectValue(a interface{})  {
	v := reflect.ValueOf(a)
	fmt.Println(v)

	value := v.Kind()
	fmt.Println(value)

	switch value {
	case reflect.Float64:
		fmt.Println("a是：", v.Float())
	}
}

// 反射修改值信息

func reflectSetValue(a interface{})  {
	v := reflect.ValueOf(a)
	k := v.Kind()

	switch k {
	case reflect.Float64:
		v.SetFloat(6.5)
		fmt.Println(a)
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(6.6)
		fmt.Println(v.Elem().Float())
		fmt.Println(v.Pointer())
	}


}

func main() {
	var x float64 = 3.4
	//reflectType(x)
	//reflectValue(x)

	// 反射认为下面是指针类型，不是float类型
	reflectSetValue(&x)
	fmt.Println(x)

}