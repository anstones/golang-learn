package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 异步处理

func AsyncHandler(c *gin.Context)  {
	copyContext := c.Copy()

	go func() {
		time.Sleep(3*time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}


// 同步处理
func FuncHandler(c *gin.Context)  {
	time.Sleep(3*time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}


func main() {
	r := gin.Default()
	r.GET("/long_async", AsyncHandler)
	r.GET("/long_sync", FuncHandler)
	r.Run(":8080")
}
