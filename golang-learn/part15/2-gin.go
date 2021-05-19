package main

// 上传文件

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20

	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil{
			c.String(500, "error")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
	})
	
	// 限制文件大小和格式
	r.POST("/upload-out-size", func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		if err != nil{
			log.Printf("Error when try to get file: %v", err)
		}
		if headers.Size > 1024*1024*2{
			fmt.Println("文件太大了")
			return
		}

		if headers.Header.Get("Content-Type") != "image/png"{
			fmt.Println("只允许上传png图片")
			return
		}
		c.SaveUploadedFile(headers, "./video/"+headers.Filename)
		c.String(http.StatusOK, headers.Filename)
	})

	// 多文件上传
	r.POST("/upload-many", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil{
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		files := form.File["files"]
		for _ , file := range files{
			if err := c.SaveUploadedFile(file, file.Filename); err != nil{
				c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload ok %d files", len(files)))
	})
	
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
	r.Run(":8080")
}
