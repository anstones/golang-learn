package main

import (
	"fmt"
	"net"
)

// tcp粘包问题

func nianbaoClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}

	defer conn.Close()
	for i := 0; i < 20; i++{
		msg := "Hello, Hello. How are you?"
		_,_ = conn.Write([]byte(msg))
	}
}

