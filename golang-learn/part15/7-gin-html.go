package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("html/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index1.html", gin.H{"title": "我是测试", "address": "www.5lmh.com"})
	})
	// 重定向
	r.GET("/redirect", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "http://www.5lmh.com")
	})
	r.Run(":8080")
}