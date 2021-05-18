package main

import (
	"fmt"
	"golang-learn/part13/TestReflect"
	"io/ioutil"
)

func parseFile(fileNme string)  {
	data, err  := ioutil.ReadFile(fileNme)
	if err != nil{
		fmt.Println(err)
		return
	}

	var cong TestReflect.Config

	err = TestReflect.UnMarshal(data, &cong)
	if err != nil{
		return
	}
	fmt.Printf("反序列化成功  conf: %#v\n  port: %#v\n", cong, cong.ServerConf.Port)
}

func parseFile2(filename string)  {
	// 有一些假数据
	var conf TestReflect.Config
	conf.ServerConf.Ip="127.0.0.1"
	conf.ServerConf.Port=8000
	conf.MysqlConf.Port=9000
	err := TestReflect.MarshalFile(filename,conf)
	if err != nil{
		fmt.Println(err)
		return
	}
}

func main() {
	//parseFile("F:/gospace/src/golang-learn/part13/TestReflect/conf.ini")
	parseFile2("F:/gospace/src/golang-learn/part13/TestReflect/conf.ini")
}