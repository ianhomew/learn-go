package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 實作一個 Hero 結構體切片 並進行排序 sort.Sort(data interface)

type Hero struct {
	Name string
	Age  int
}

// Hero 結構體切片
type HeroSlice []Hero

// 實現接口
//type Interface interface {
//	// Len is the number of elements in the collection.
//	Len() int
//	// Less reports whether the element with
//	// index i should sort before the element with index j.
//	Less(i, j int) bool
//	// Swap swaps the elements with indexes i and j.
//	Swap(i, j int)
//}
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 方法決定排序方式
// 按照 Age 從小到大
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
}
func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

func main() {

	//var intSlice = []int{0, -1, 10, 7, 90}
	//sort.Ints(intSlice)
	//fmt.Println(intSlice)

	rand.Seed(time.Now().UnixNano())

	var heroes HeroSlice
	for i := 0; i < 5; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄_%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	fmt.Println("before sort")
	for _, v := range heroes {
		fmt.Println(v)
	}

	sort.Sort(heroes)

	fmt.Println("after sort")
	for _, v := range heroes {
		fmt.Println(v)
	}

}
