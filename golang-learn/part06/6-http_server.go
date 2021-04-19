package main

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request){
	_ = r.ParseForm() // 解析参数，默认是不会解析的
	fmt.Println("method", r.Method)
	fmt.Println("url", r.URL.Path)
	fmt.Println("header:", r.Header)
	fmt.Println("body", r.Body)
	_, _ = w.Write([]byte("www.51mh.com"))
}

func main()  {
	http.HandleFunc(
		"/go",
		Handler,
	)
	_ = http.ListenAndServe("127.0.0.1:8000", nil)
}
