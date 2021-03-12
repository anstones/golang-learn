package main

import (
	"fmt"
	"os"
	"time"
)

// 自定义异常

type PathError struct {
	path, op, createTime, message string
}


func (p *PathError) Error() string{
	return fmt.Sprintf("path=%s \nop=%s \ncreateTime=%s \nmessage=%s", p.path,
		p.op, p.createTime, p.message)
}


func Open(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil{
		return &PathError{
			path: fileName,
			op:"read",
			message: err.Error(),
			createTime: fmt.Sprintf("%v", time.Now()),
		}
	}
	defer file.Close()
	return nil
}

func main() {
	err := Open("/Users/5lmh/Desktop/go/src/test.txt")

	switch value := err.(type) {
	case *PathError:
		fmt.Println("get path error,", value)
	default:

	}
}