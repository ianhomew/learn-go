package main

import "fmt"

func main() {

	// 使用 select 解決管道阻塞問題

	// 1. 定義一個管道 10 個數據 int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	// 2. 定義一個管道 5 個數據 string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello " + fmt.Sprintf("%d", i)
	}

	// 傳統的方法在 loop 管道時 如果沒有先關閉 會阻塞而導致 dead lock
	// 問題： 在實際開發中 不好確定什麼時候該關閉管道
	// 可以使用 select 方式解決

	for {
		select {
		case v := <-intChan: // 注意 這裡如果管道一直沒有關閉 也不會一直阻塞導致 dead lock 會自動往下一個 case 匹配
			fmt.Println("從 intChan 讀取了數據: ", v)
		case v := <-stringChan:
			fmt.Println("從 stringChan 讀取了數據: ", v)
		default:
			fmt.Println("都取不到了 不玩了")
			return
		}
	}

}
