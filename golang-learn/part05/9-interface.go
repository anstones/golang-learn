package main

// 空接口的应用  接口类型断言


import (
	"bytes"
	"fmt"
	"io"
	"os"
)


func testType(){
	var x interface{}
	x = "122222"
	v, ok := x.(string)
	if ok{
		fmt.Println(v)
	}else{
		fmt.Println("类型断言失败")
	}
}


func justifyType(x interface{}){
	switch v:=x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}



func main() {

	var x string  = "123"
	fmt.Printf("type %T, value %v\n", x, x)

	testType()

	var w io.Writer
	w = os.Stdout
	w = new(bytes.Buffer)
	w = nil
	fmt.Println(w)

	var d = "stone"
	justifyType(d)

}