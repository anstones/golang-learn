package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://127.0.0.1:8000/go")
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	buf := make([]byte, 1024)
	for {
		res, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF{
			fmt.Println(err)
			return
		}else {
			fmt.Println("读取完毕")
			fmt.Println(res)
			break
		}

	}
}