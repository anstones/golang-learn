package proto

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn)  {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		msg , err := Decode(reader)
		if err != io.EOF{
			return
		}
		if err != nil{
			fmt.Println("decode msg failed, err:", err)
			return
		}

		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:")

	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}

	defer listen.Close()
	for{
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed, err:", err)
			continue
		}
		process(conn)
	}


}