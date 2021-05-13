package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	reQQEmail = `(\d+)@qq.com`
)

func ErrorHandler(err error, why string)  {
	if err != nil{
		fmt.Println(why, err)
	}
}


func GetEmail()  {
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	ErrorHandler(err, "http.Get url")
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	ErrorHandler(err, "ioutil.ReadAll")

	//字节转字符串
	page_str := string(pageBytes)

	re := regexp.MustCompile(reQQEmail)

	results := re.FindAllStringSubmatch(page_str, -1)

	for _, result := range results{
		fmt.Println("email:", result[0])
		fmt.Println("qq:", result[1])
	}

}

func main() {
	GetEmail()
}