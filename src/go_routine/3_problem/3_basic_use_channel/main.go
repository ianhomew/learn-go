package main

import "fmt"

// 使用鎖的問題
// 主線程在等待所有 go routine 全部完成 但是等待多久只能推估
// 如果推估的時間比較短 則部分 go routine 可能還處於工作狀態 這時會因為主線程的結束 造成協程的銷毀
// 通過全局變數添加鎖實現通訊 會不利於多個協程對全局變數的讀寫操作 (意思是 工程師還要知道在哪裡加鎖)
//
// 基於以上 我們需要 channel

// 管道 channel 介紹:
// channel 的本質 就是數據結構 - 隊列 queue
// 先進先出
// 線程安全(編譯器底層維護) 不需要加鎖
// channel 是有類型的 string 類型的 channel 只能存放 string 類型的資料
// 多個協程操作同一個管道時 不會發生資源競爭問題

// channel 是引用類型
// channel 必須初始化才能寫數據 (make後使用)

func main() {

	var intChan chan int = make(chan int, 3)

	// 是記憶體位址 所以是 引用類型
	fmt.Printf("intChan value = %v, intChan 本身的位址 = %p\n", intChan, &intChan)

	// write data to channel
	intChan <- 10

	num := 211
	intChan <- num

	// 注意: 往管道寫資料的時候 不能超過 cap 容量限制
	// fatal error: all goroutines are asleep - deadlock!
	//intChan <- 11
	//intChan <- 12

	// 看管道的長度和 cap容量
	// 管道的 cap 不會自動長大
	// channel len = 2, cap = 3
	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))

	// 從管道中讀取數據
	num2 := <-intChan
	fmt.Println("num2=", num2)
	fmt.Printf("channel len = %v, cap = %v\n", len(intChan), cap(intChan))

	// 在沒有使用協程的情況下 如果管道資料已經全部取出了 就會出錯
	// 因為主線程會無限等待直到有數據寫入管道 但是根本不可能有數據寫入管道了 所以會出錯
	// fatal error: all goroutines are asleep - deadlock!
	//fmt.Println(<- intChan)
	//fmt.Println(<- intChan)
	//fmt.Println(<- intChan)

}
