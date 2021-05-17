package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// input
func bufioDemo()  {
	reader  := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	text,_ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Printf("%v\n", text)
}


func main() {
	// Print
	//func Print(a ...interface{}) (n int, err error)
	//func Printf(format string, a ...interface{}) (n int, err error)
	//func Println(a ...interface{}) (n int, err error)

	// Fprint
	//fmt.Fprintln(os.Stdout, "项标准输出写入类容")
	//
	//file_obj, err := os.OpenFile("xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	//if err != nil{
	//	fmt.Println("打开文件出错，", err)
	//	return
	//}
	//name := "枯藤"
	//// 向打开的文件句柄中写入内容
	//fmt.Fprintf(file_obj, "往文件中写如信息：%s", name)
	//
	//// Sprint
	//s1 := fmt.Sprint("ssssssss")
	//name = "ssssssssss"
	//age := 18
	//s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	//s3 := fmt.Sprintln("qqqqqqqq")
	//fmt.Println(s1+s2+s3)
	//
	//
	//// Errorf
	//err = fmt.Errorf("error")
	//fmt.Println(err)

	//var (
	//	name string
	//	age int
	//	married bool
	//)
	//
	//fmt.Scan("1:%s 2:%d 3:%t", &name, &age, &married)
	//
	//fmt.Printf("扫描结果 name:%s age:%d married:%t \n", name, age, married)


	bufioDemo()

}