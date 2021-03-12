package main

import (
	"fmt"
	"net/http"
)

func do() error {
	res, err := http.Get("http://www.google.com")
	if res != nil{
		defer res.Body.Close()  // 总是在一次成功的资源分配下面使用 defer ，对于这种情况来说意味着：当且仅当 http.Get 成功执行时才使用 defer
	}
	if err != nil{
		return err
	}
	// ----code----
	fmt.Println(res)
	return nil
}

func main() {
	do()
}