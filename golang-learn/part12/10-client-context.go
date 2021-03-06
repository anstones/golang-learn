package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

//调用服务端API时如何在客户端实现超时控制？

type respData struct {
	resp *http.Response
	err error
}

func doCall(ctx context.Context)  {
	transport := http.Transport{
		// 请求频繁可定义全局的client对象并启用长链接
		// 请求不频繁使用短链接
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}

	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("get", "http://127.0.0.1:9090", nil)
	if err != nil{
		fmt.Printf("new requestg failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx)  // 使用带超时的ctx创建一个新的client request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err: err,
		}
		respChan <- rd
		wg.Done()
	}()

	select {
	case <- ctx.Done():
		fmt.Println("call api timeout")
	case result := <- respChan:
		fmt.Println("call server api success")
		if result.err != nil{
			fmt.Printf("call server api failed, err:%v\n", result.err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Println(string(data))
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*100)
	defer cancel()  // 调用cancel释放子goroutine资源
	doCall(ctx)
}