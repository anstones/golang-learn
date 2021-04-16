package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp 服务端

func process(conn net.Conn)  {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [1024]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil{
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		_, _ = conn.Write([]byte(recvStr)) // 回显

	}

}

func main() {
	listion, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil{
		fmt.Println("listioned failed, err:", err)
		return
	}
	for {
		conn, err := listion.Accept() // 建立连接
		if err != nil{
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)  // 启动一个goroutine处理连接
	}
}