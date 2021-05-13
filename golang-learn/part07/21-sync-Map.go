package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	m[key] = value
}

// 并发不安全
//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go func(n int) {
//			key := strconv.Itoa(n)
//			set(key, n)
//			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}


var mm = sync.Map{}

func setValue (n int, wgs *sync.WaitGroup) {
	key := strconv.Itoa(n)
	mm.Store(key, n)
	value,_ := mm.Load(key)
	fmt.Printf("k=:%v,v:=%v\n", key, value)
	wgs.Done()
}


func main() {
	wgs := sync.WaitGroup{}
	for i := 0; i<20;i++{
		wgs.Add(1)
		go setValue(i, &wgs)
	}
	wgs.Wait()
}
