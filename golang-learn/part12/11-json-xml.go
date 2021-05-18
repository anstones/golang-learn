package main

import (
	"encoding/json"
	"fmt"
)

// struct->json

type Person struct {
	Name string
	Hobby string
}


func structToJsonDemo()  {
	p := Person{
		"5lmh.com",
		"女",
	}
	b, err := json.Marshal(p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))

	// 格式化输出
	b, err = json.MarshalIndent(p, "", "   ")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

type PersonMan struct {
	Name string `json:"_"`
	Hobby string `json:"hobby"`
}

func mapToJsonDemo(){
	students := make(map[string]interface{})
	students["name"] = "5lmh.com"
	students["age"] = 10
	students["sex"] = "man"
	b, err := json.Marshal(students)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

func jsonToStructDemo()  {
	jsonStr := []byte(`{"Name":"5lmh.com","Hobby":"女"}`)
	var p Person
	err := json.Unmarshal(jsonStr, &p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(p)
}

func jsonToInterfaceDemo()  {
	jsonStr := []byte(`{"Name":"5lmh.com","Hobby":"女"}`)

	var in interface{}

	err := json.Unmarshal(jsonStr, &in)
	if err != nil{
		fmt.Println(err)
	}
	// 自动转到map
	fmt.Println(in)

	m := in.(map[string]interface{})

	for k, v := range m{
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "是float64类型", vv)
		case string:
			fmt.Println(k, "是string类型", vv)
		default:
			fmt.Println("其他")
		}
	}
}



func main() {
	//structToJsonDemo()
	//mapToJsonDemo()
	//jsonToStructDemo()
	jsonToInterfaceDemo()
}