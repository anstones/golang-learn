package main

import "fmt"

type people interface {
	Speak(string) string
}


type Student8 struct {}


func (s Student8) Speak(think string) (talk string){
	if think == "SB"{
		talk = "你是个大帅比"
	}else{
		talk = "你好"
	}

	return
}

func main()  {
	var peo people = &Student8{}

	think := "bitch"

	fmt.Println(peo.Speak(think))
}