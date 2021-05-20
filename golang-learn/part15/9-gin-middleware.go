package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func middle(c *gin.Context)  {
	t := time.Now()
	fmt.Println("中间件开始执行了")
	c.Set("request", "中间件")
	status := c.Writer.Status()
	fmt.Println("中间件执行完毕", status)
	t2 := time.Since(t)
	fmt.Println("time,", t2)
}


func MiddleWare()gin.HandlerFunc  {
	//return middle
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行了")
		c.Set("request", "中间件")
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("time,", t2)
	}
}

func main() {
	r := gin.Default()

	// 全局中间件
	r.Use(MiddleWare())
	r.GET("ce", func(context *gin.Context) {
		req, _ := context.Get("request")
		fmt.Println("request", req)
		context.JSON(200, gin.H{"request": req})
	})
	
	// 局部中间件
	r.GET("ce1", MiddleWare(), func(c *gin.Context) {
		// 取值
		req, _ := c.Get("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{"request": req})
	})
	

	r.Run(":8080")
}