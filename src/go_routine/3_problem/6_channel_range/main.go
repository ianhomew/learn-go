package main

import "fmt"

// 使用 for range
// 注意:
// 使用 for range 時 如果管道沒有關閉 則會出現 deadlock錯誤
// (上面這很好用 可以用管道避免之前的無法預估時間的問題 只需要關閉管道 主線程就會知道協程已經結束)
// 使用 for range 時 如果管道已經關閉 則會正常 loop, loop 結束後 就會退出 for range

func main() {

	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放一百個數字到管道內
	}

	// loop
	// 「保證錯誤」 因為每取一個 len 就會減一
	//for i := 0; i < len(intChan2); i++ {
	//
	//}

	// 這樣不會報錯
	//var len1 = len(intChan2)
	//for i := 0; i < len1; i++ {
	//	fmt.Println(<-intChan2)
	//}

	// loop 前 要關閉 否則報錯
	// fatal error: all goroutines are asleep - deadlock!
	close(intChan2)

	//loop
	for v := range intChan2 {
		fmt.Println(v)
	}

}
