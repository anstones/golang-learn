package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 文件相关操作
func openFileDemo()  {
	file, err := os.Open("xx.txt")

	if err != err{
		fmt.Println("open file failed!, err:", err)
		return
	}
	file.Close()
}

func WriteFileDemo()  {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i <10; i++{
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
	}
}

func ReadFileDemo()  {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("open file err :", err)
		return
	}
	defer file.Close()

	// 定义文件接收的数组
	var buf [256]byte
	var content []byte

	for {
		n, err := file.Read(buf[:])
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			return
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

func copyFileDemo()  {
	file, err := os.Open("test.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	newFile, err := os.Create("test_bak.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer newFile.Close()

	var buf [1024]byte
	for {
		n, err := file.Read(buf[:])
		if err == io.EOF{
			fmt.Println("读取完毕")
			break
		}
		if err != nil{
			fmt.Println(err)
			return
		}
		newFile.Write(buf[:n])
	}
}


// bufio 读文件

func BufioWriteFile()  {
	file, err := os.Create("xxx.txt")
	if err != nil{
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i :=0; i<10;i++{
		writer.WriteString("qqqqqq")
	}
	// 强制写入
	writer.Flush()
}

func BufioReadFile()  {
	file, err  := os.Open("xxx.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(string(line))
	}
}

//  ioutil工具包

func ioutilWriteFile()  {
	err := ioutil.WriteFile("xxxx.txt", []byte("www.51mh.com"), 0666)
	if err != nil{
		fmt.Println(err)
		return
	}
}


func ioutilReadFile(){
	content, err := ioutil.ReadFile("xxxx.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

// 实现cat命令

func Cat(r *bufio.Reader)  {
	for {
		buf , err := r.ReadBytes('\n')
		if err == io.EOF{
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}


func CatReadDemo()  {
	flag.Parse()  // 解析命令行参数
	if flag.NArg() == 0{
		Cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容打印到终端
	for i := 0; i<flag.NArg(); i++{
		f , err := os.Open(flag.Arg(i))
		if err != nil{
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		Cat(bufio.NewReader(f))
	}
}


func main() {
	//openFileDemo()
	//WriteFileDemo()
	//ReadFileDemo()
	//copyFileDemo()
	//BufioWriteFile()
	//BufioReadFile()
	//ioutilWriteFile()
	ioutilReadFile()
}
