package main

import "fmt"

// 結構體是值類型
// slice, map, pointer 默認值是 nil 表示尚未分配記憶體空間
// 若要使用 slice or map1 需要 make
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {

	var p1 Person

	if p1.ptr == nil {
		fmt.Println("p1.ptr is nil")
	}
	if p1.slice == nil {
		fmt.Println("p1.slice is nil")
	}
	if p1.map1 == nil {
		fmt.Println("p1.map1 is nil")
	}

	// 要先 make 否則報錯
	p1.slice = make([]int, 2)
	p1.slice[0] = 100
	fmt.Println(p1)

}
