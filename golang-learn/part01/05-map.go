package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func map1() {
	m := make(map[string]int, 8)
	m["张三"] = 6
	m["小明"] = 20
	fmt.Println(m)
	fmt.Println(m["小明"])
	fmt.Printf("type of m:%T\n", m)

	n := map[string]string{
		"username": "pprof.cn",
		"password": "123456",
	}
	fmt.Println(n)
}

// 判断map中键值是否存在
func map2() {
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 1
	scoreMap["小明"] = 2

	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

// 按照指定顺序遍历map
func map3() {
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// 元素为map的切片
func map4() {
	var mapslice = make([]map[string]string, 3)

	for index, value := range mapslice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}

	fmt.Println("affter init")

	mapslice[0] = make(map[string]string, 30)
	mapslice[0]["name"] = "小明"
	mapslice[0]["password"] = "123456"
	mapslice[0]["address"] = "红旗大街"

	for index, value := range mapslice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

// 值为切片类型的map
func map5() {
	var slicemap = make(map[string][]string, 3)

	fmt.Println(slicemap)

	fmt.Println("affter init")

	key := "中国"
	value, ok := slicemap[key]
	if !ok {
		value = make([]string, 0, 2)

	}
	value = append(value, "北京", "上海")
	slicemap[key] = value
	fmt.Println(slicemap)
}

// 当往map中存储一个kv对时，通过k获取hash值，
// hash值的低八位和bucket数组长度取余，定位到在数组中的那个下标，
// hash值的高八位存储在bucket中的tophash中，用来快速判断key是否存在，
// key和value的具体值则通过指针运算存储，当一个bucket满时，通过overfolw指针链接到下一个bucket。

func main() {
	//map1()
	//map2()
	//map3()
	//map4()
	map5()
}
