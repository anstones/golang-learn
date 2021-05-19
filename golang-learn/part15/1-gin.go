package main

import (
"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/:action", func(c *gin.Context) {
		// 解析参数
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		// DefaultQuery() 设置默认参数
		// PostForm 表单参数
		c.String(http.StatusOK, name + " is "+action)
	})
	//监听端口默认为8080
	r.Run(":8000")
}
