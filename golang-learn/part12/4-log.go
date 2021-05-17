package main

import (
	"fmt"
	"log"
	"os"
)

func init()  {
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile|log.Lmicroseconds|log.Ldate)
}

func main() {
	//logger := log.New(os.Stdout, "[<New>]", log.Lshortfile|log.Ldate|log.Ltime)
	log.Println("这是自定义的logger记录的日志。")
	log.Fatalln("这是一条会触发fatal的日志。")
}