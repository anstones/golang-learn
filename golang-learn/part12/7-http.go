package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// 内置net/http

func httpGetDemo()  {
	resp, err := http.Get("https://www.5lmh.com/")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err :=   ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func argsHttpGetDemo()  {
	apiUrl := "http://127.0.0.1:9090/get"
	// url params
	data := url.Values{}
	data.Set("name", "kk")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil{
		fmt.Println(err)
		return
	}
	u.RawQuery = data.Encode() // url encode
	fmt.Println(u.String())
	
	resp, err := http.Get(u.String())
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func httpPostDemo()  {
	apiUrl := "http://127.0.0.1:9090/post"
	contentType := "application/json"
	data := `{"name":"st", "age":22}`

	resp ,err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil{
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}


func sayHello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "hello golang")
}


func httpServerDemo()  {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil{
		fmt.Println(err)
		return
	}
}



func main() {
	//httpGetDemo()
	//httpPostDemo()
	httpServerDemo()
}



