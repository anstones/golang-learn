package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Passwd string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

//客户端传参，后端接收并解析到结构体
func JsonHandler(c *gin.Context)  {
	var json Login
	err := c.ShouldBindJSON(&json)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if json.User != "root" || json.Passwd != "admin"{
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

// 表单数据解析和绑定
func FormHandler(c *gin.Context)  {
	var form Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&form); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if form.User != "root" || form.Passwd != "admin"{
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}


// URI数据解析和绑定
func UriHandler(c *gin.Context)  {
	var login Login
	if err := c.ShouldBindUri(&login); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if login.User != "root" || login.Passwd != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}


func main() {
	r := gin.Default()
	r.POST("loginJSon", JsonHandler)
	r.POST("loginForm", FormHandler)
	r.POST("loginUri", UriHandler)
	r.Run(":8080")
}
