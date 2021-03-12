package main

import "fmt"

func traversalString() {
	s := "pprof.cn博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}

	fmt.Println()
	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}

func changeString() {

	// 要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
	// 无论哪种转换，都会重新分配内存，并复制字节数组。

	s1 := "hello嘿嘿"
	byteS1 := []byte(s1)
	fmt.Println("type ,", byteS1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runes2 := []rune(s2)
	runes2[0] = '狗'
	fmt.Println(string(runes2))
}

func main() {
	//traversalString()
	changeString()
}
