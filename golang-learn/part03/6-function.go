package main

import "fmt"

type TestDefer struct {
	name string
}

func (t *TestDefer)closed()  {
	fmt.Println(t.name, " closed")
}


func test_defer_3(){
	ts := []TestDefer{{"a"}, {"b"}, {"c"}}
	for _,i := range ts{
		//defer i.closed()
		/*
		defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。但是并没有说struct这里的this指针如何处理，
		通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。
		*/
		t:= i
		defer t.closed()
	}
}

func main() {
	test_defer_3()
}
