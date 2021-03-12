package main

import (
	"fmt"
	"sort"
)

func main()  {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"

	sli := []int{}
	for k, _ :=range map1{
		sli = append(sli, k)
	}
	sort.Ints(sli)
	fmt.Println(sli)
	for i := 0; i<len(sli); i++{
		fmt.Println(map1[sli[i]])
	}

}

