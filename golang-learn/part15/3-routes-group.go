
package main

// routes 格式

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func login(c *gin.Context)  {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context)  {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}


func main() {
	r := gin.Default()

	v1 := r.Group("/v1")
	v2 := r.Group("/v2")
	{
		v1.GET("login", login)
		v1.GET("submit", submit)
		v2.POST("login", login)
		v2.POST("submit", submit)
	}
	r.Run(":8080")
}
