/*
 * @Author: your name
 * @Date: 2021-03-08 08:32:18
 * @LastEditTime: 2021-03-08 08:42:30
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gospace\src\golang-learn\part01\12-struct.go
 */

package main

import "fmt"

type studentmore struct {
	id   int
	name string
	age  int
}

func demo(ce []studentmore){
	//切片是引用传递，是可以改变值的
	ce[1].age = 999
}

func delete_map_struct(ce map[int]studentmore)  {
	delete(ce, 2)
}



func main() {
	var ce []studentmore
	ce = []studentmore{
		{1, "s", 22},
		{2, "t", 33},
	}
	fmt.Println(ce)
	demo(ce)
	fmt.Println(ce)


	cl := make(map[int]studentmore)
	cl[1] = studentmore{1,"b", 22}
	cl[2] = studentmore{2,"c", 88}
	fmt.Println(cl)
	delete_map_struct(cl)
	fmt.Println(cl)
}
