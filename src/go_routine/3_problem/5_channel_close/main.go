package main

import "fmt"

// 使用內建的 close() 可以關閉管道 關閉之後 只能讀 不能寫

func main() {

	var intChan = make(chan int, 3)

	intChan <- 3
	intChan <- 100

	close(intChan)
	// panic: send on closed channel
	//intChan <- 200
	// 可以讀
	fmt.Println(<-intChan)

}
