package main

import (
	"fmt"
	"reflect"
)

// 结构体与反射

type UserMan struct {
	Id int
	Name string
	Age int
}

func (u UserMan)Hello()  {
	fmt.Println("hello")
}

func Poni(o interface{})  {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())

	value := reflect.ValueOf(o)
	fmt.Println(value)

	for i := 0; i<t.NumField();i++{
		k:= t.Field(i)
		fmt.Printf("%s : %v\n", k.Name, k.Type)

		v := value.Field(i).Interface()
		fmt.Printf("%v:%v\n", k.Name, v)
	}

	fmt.Println("==============获取方法++++++++++++++++++++++")
	for i := 0; i<t.NumMethod();i++{
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}

}

// 修该结构体的值
func setValueDemo(i interface{})  {
	v := reflect.ValueOf(i)

	// 获取指针元素
	value := v.Elem()
	fmt.Println(value)

	f := value.FieldByName("Name")
	fmt.Println(f, f.Kind())
	if f.Kind() == reflect.String{
		f.SetString("st")
	}

}


// 调用方法

func (u UserMan)HelloFunc(name string)  {
	fmt.Println("hello: ", name)
}

func reflectCallStructMethod(i interface{})  {
	//v := reflect.TypeOf(i)
	k := reflect.ValueOf(i)
	// 函数名首字母必须大写
	n := k.MethodByName("HelloFunc")
	args := []reflect.Value{reflect.ValueOf("6666")}
	n.Call(args)
}






func main() {
	u := UserMan{1, "zs", 20}
	//Poni(u)

	//setValueDemo(&u)
	//fmt.Println(u)

	//reflectCallStructMethod(u)
}