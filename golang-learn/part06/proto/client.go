package proto

import (
	"fmt"
	"net"
)

func main() {
	conn , err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i<20; i++{
		mas := `Hello, Hello. How are you?`
		data, err := Encode(mas)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}