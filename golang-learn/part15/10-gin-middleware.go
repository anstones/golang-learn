package main

import (
	"fmt"
	_ "github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"time"
)

func MyTimeMiddleWare(c *gin.Context)  {
	st := time.Now()
	c.Next()
	since := time.Since(st)
	fmt.Println("程序用时：", since)
}


func shopIndexHandler(c *gin.Context)  {
	time.Sleep(5*time.Second)
	t, _  := c.Get("t")
	c.JSON(200, gin.H{"用时":t})
}

func shopHomeHandler(c *gin.Context)  {
	time.Sleep(3 * time.Second)
	t, _  := c.Get("t")
	c.JSON(200, gin.H{"用时":t})
}

func main() {
	r := gin.Default()
	r.Use(MyTimeMiddleWare)

	shopingGrep := r.Group("/shopping")
	{
		shopingGrep.GET("/index", shopIndexHandler)
		shopingGrep.GET("/home", shopHomeHandler)
	}
	r.Run(":8080")
}