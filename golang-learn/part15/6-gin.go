package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
)

// json、结构体、XML、YAML类似于java的properties、ProtoBuf

func ReturnJson(c *gin.Context)  {
	c.JSON(200, gin.H{"message": "someJSON", "status": 200})
}

func ReturnStruct(c *gin.Context)  {
	var msg struct{
		Name string
		Message string
		Number int
	}
	msg.Name = "root"
	msg.Message = "msg"
	msg.Number = 123
	c.JSON(200, msg)
}

func ReturnXml(c *gin.Context)  {
	c.XML(200, gin.H{"message":"abc"})
}

func ReturnYaml(c *gin.Context)  {
	c.YAML(200, gin.H{"name": "zhangsan"})
}

func ReturnProtoBuf(c *gin.Context)  {
	rep := []int64{int64(1), int64(2)}
	label := "label"
	data := &protoexample.Test{
		Label: &label,
		Reps: rep,
	}
	c.ProtoBuf(200, data)
}

func main() {
	r := gin.Default()
	// json 响应
	r.GET("/soneJSON", ReturnJson)
	r.GET("/someStruct", ReturnStruct)
	r.GET("/someXML", ReturnXml)
	r.GET("/someYAML", ReturnYaml)
	r.GET("/someProtoBuf", ReturnProtoBuf)
	r.Run(":8080")
}
