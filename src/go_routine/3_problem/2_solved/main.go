package main

import (
	"fmt"
	"sync"
	"time"
)

// 寫一個 1 - 20 的各個數的階層 並且把階層放入到 map 內

var (
	myMap = make(map[int]int, 10)
	// 宣告一個全局的互斥鎖

	// lock 是一個全局互斥鎖
	// Mutex 是互斥的意思
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()
	myMap[n] = res // concurrent map writes here was gone.
	lock.Unlock()
}

func main() {
	// 1. 寫一個函數 負責計算每個數的階層 並放入到 map 內
	// 2. 啟動多個協程 統一將階層計算的結果 放入 map 內
	// 3. 每個協程都要訪問 map 所以 map 應該是 global

	for i := 1; i <= 20; i++ {
		go test(i)
	}

	// 非常大的問題是 我根本不知道這邊要設定幾秒
	// 好的硬體比較少秒 壞的硬體會等比較多秒
	// 準確等多少秒無法得知
	// 等待的目的是要等 go test()做完計算 否則程式就直接往下跑了
	time.Sleep(5 * time.Second)

	lock.Lock()
	fmt.Println("len = ", len(myMap))
	for i, v := range myMap {
		fmt.Printf(" %v! = %v\n", i, v)
	}
	lock.Unlock()

	// for range 如果沒加鎖 會出錯 why?
	// 在實際程式運行中 還是可能造成 資源競爭的問題
	// 因為人知道 5 秒可以做完
	// 但是「程式」不知道
	// fatal error: concurrent map iteration and map write

}
