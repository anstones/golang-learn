package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 并发爬虫

var (
	chanImageUrl chan string
	wg sync.WaitGroup
	chanTask chan string
	reimg = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)


func HanleError(err error, why string)  {
	if err != nil{
		fmt.Println(why, err)
	}
}

func GetPagestr(url string) (pagestr string) {
	resp, err := http.Get(url)
	HanleError(err, "http.Get url")
	defer resp.Body.Close()

	pageBytes, err := ioutil.ReadAll(resp.Body)
	HanleError(err, "ioutil.ReadAll")

	pagestr = string(pageBytes)
	return
}



func GetFilenameFromUrl(url string)(fileName string)  {
	lastIndex := strings.LastIndex(url, "/")
	fileName = url[lastIndex+1:]

	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	fileName = timePrefix + "_" +fileName
	return
}


func DownloadFile(url string, fileName string) (ok bool)  {
	resp, err := http.Get(url)
	HanleError(err, "http.get.url")
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	HanleError(err, "resp.body")

	fileName = "F:\\gospace\\src\\golang-learn\\part08" + fileName

	err = ioutil.WriteFile(fileName, bytes, 0666)

	if err != nil{
		return false
	}else {
		return true
	}

}

func DownloadImg()  {
	for url := range chanImageUrl{
		fileName := GetFilenameFromUrl(url)
		ok := DownloadFile(url, fileName)
		if ok {
			fmt.Printf("%s 下载成功\n", fileName)
		} else {
			fmt.Printf("%s 下载失败\n", fileName)
		}
	}
	wg.Done()
}

func getImgs(url string) (urls []string)  {
	pageStr := GetPagestr(url)
	re := regexp.MustCompile(reimg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))

	for _, result := range results{
		url := result[0]
		urls = append(urls, url)
	}
	return
}

func getImgUrls(url string)  {
	urls := getImgs(url)
	for _, url := range urls{
		chanImageUrl <- url
	}
	chanTask <- url
	wg.Done()
}

func checkOk()  {
	var count int
	for {
		url := <- chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count ++
		if count == 26{
			close(chanImageUrl)
			break
		}
	}
	wg.Done()
}



func main() {
	chanImageUrl = make(chan string, 1000000)
	chanTask = make(chan string, 26)

	for i:=1; i<27;i++{
		wg.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E5%8F%AF%E7%88%B1/" + strconv.Itoa(i) + ".html")
	}

	wg.Add(1)
	go checkOk()

	for i := 0; i<5; i++{
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}

