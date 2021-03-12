/*
 * @Author: your name
 * @Date: 2021-03-08 15:12:18
 * @LastEditTime: 2021-03-08 08:23:24
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gospace\src\golang-learn\part01\11-struct.go
 */
package main

import (
	"encoding/json"
	"fmt"
)

//结构体标签（Tag）

type StudentTag struct {
	ID     int    `json:id` // 通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := StudentTag{
		ID:     1,
		Gender: "女",
		name:   "pprof",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data)
}
