package main

import (
	"fmt"
	"sort"
)

func main() {

	// map 本身是無序的
	map1 := make(map[int]int, 10)

	map1[10] = 100
	map1[1] = 40
	map1[4] = 30
	map1[8] = 90

	fmt.Println(" fmt.print 顺序印出來, range 確實是無序的")
	for k, v := range map1 {
		fmt.Printf("map1[%d] = %v\n", k, v)
	}

	//fmt.Println(map1)

	// 1. 先將 map 的 key 放到切片內
	var keys []int
	for k := range map1 {
		keys = append(keys, k)
	}
	// 2. 對切片進行排序
	sort.Ints(keys)
	fmt.Println(keys)
	// 3. 循環切片 按照 key 輸出 map 的值
	for _, v := range keys {
		fmt.Printf("map1[%d] = %v\n", v, map1[v])
	}

}
