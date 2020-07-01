package main

import "fmt"

// 只能寫 main 內的 ch 可以是雙向的
func send(ch chan<- int, exitChan chan struct{}) {

}

// 只能讀 main 內的 ch 可以是雙向的
func recv(ch <-chan int, exitChan chan struct{}) {

}

func main() {

	// 管道可以聲明為只讀或者只寫

	// 1. 默認 雙向
	var chan1 chan int

	// 2. 只寫
	var chan2 chan<- int
	chan2 = make(chan int, 10)
	chan2 <- 20
	// 不能讀
	//fmt.Println(<-chan2)

	// 3. 只讀
	var chan3 <-chan int
	chan3 = make(chan int, 10)
	//chan3 <- 10 // error
	<-chan3

}
