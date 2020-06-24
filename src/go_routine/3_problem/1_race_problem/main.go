package main

import (
	"fmt"
	"time"
)

// 寫一個 1 - 20 的各個數的階層 並且把階層放入到 map 內

var (
	myMap = make(map[int]int, 10)
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	myMap[n] = res // concurrent map writes here
}

func main() {
	// 1. 寫一個函數 負責計算每個數的階層 並放入到 map 內
	// 2. 啟動多個協程 統一將階層計算的結果 放入 map 內
	// 3. 每個協程都要訪問 map 所以 map 應該是 global

	for i := 1; i <= 20; i++ {
		// 因為多個協程 都在操作 myMap
		go test(i)
	}

	time.Sleep(10 * time.Second)

	for i, v := range myMap {
		fmt.Printf(" %d! = %d\n", i, v)
	}

}
