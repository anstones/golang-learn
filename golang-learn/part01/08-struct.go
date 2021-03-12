/*
 * @Author: your name
 * @Date: 2021-03-08 15:12:18
 * @LastEditTime: 2021-03-08 08:31:00
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: \gospace\src\golang-learn\part01\08-struct.go
 */
package main

import "fmt"

type Person1 struct {
	string
	int
}

func struct02() {
	// 结构体的匿名字段
	p02 := Person1{
		"pprof.cn",
		18,
	}

	fmt.Printf("%#v\n", p02)
	// 匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
	fmt.Println(p02.string, p02.int)
}

func main() {
	struct02()
}
