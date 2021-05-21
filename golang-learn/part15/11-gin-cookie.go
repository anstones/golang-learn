package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CookieMiddleWare() (gin.HandlerFunc) {
	return func(context *gin.Context) {
		cookie, err := context.Cookie("abc")
		if err == nil{
			if cookie == "123"{
				context.Next()
				return
			}
		}
		context.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		context.Abort()
		return
	}
}



func main() {
	r := gin.Default()
	r.GET("/login", func(context *gin.Context) {
		context.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		//返回信息
		context.String(200, "Login success")
	})
	r.GET("/home", CookieMiddleWare(), func(context *gin.Context) {
		context.JSON(200, gin.H{"data":"home"})

	})
	r.Run(":8080")
}