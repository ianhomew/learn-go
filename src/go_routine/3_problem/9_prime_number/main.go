package main

import (
	"fmt"
	"time"
)

// 統計 1 - 8000 有哪些質數
// 統計質數的任務分配給 4 個 go routine 去完成

// 1 不是質數
func pushNumber(intChan chan int) {
	for i := 2; i <= 80; i++ {
		intChan <- i
	}
	//fmt.Println("push [", i, "] to chan")
	close(intChan)
}

func calPrimeNumber(intChan chan int, primeChan chan int, exitChan chan bool) {

	for {
		var isPrime bool
		num, ok := <-intChan
		if !ok {
			break
		}

		// 判斷是否為質數
		for i := 2; i < num; i++ {
			// 先假設是質數
			isPrime = true
			if (num % i) == 0 {
				// 不是質數
				isPrime = false
				break
			}
		}
		if isPrime {
			primeChan <- num
		}

	}
	exitChan <- true
	fmt.Println("有一個協程退出了")

}

func main() {

	var intChan = make(chan int, 8000)
	var primeChan = make(chan int, 4000)
	var exitChan = make(chan bool, 4)

	go pushNumber(intChan)
	time.Sleep(time.Second * 1)
	go calPrimeNumber(intChan, primeChan, exitChan)
	go calPrimeNumber(intChan, primeChan, exitChan)
	go calPrimeNumber(intChan, primeChan, exitChan)
	go calPrimeNumber(intChan, primeChan, exitChan)

	go func() {
		// 這樣比較厲害
		for i := 0; i < 4; i++ {
			// 這邊取不到 會等待
			<-exitChan
		}
		close(primeChan)
		close(exitChan)
	}()

	// 下面是我的寫法 不好
	//var count int
	//for {
	//	v, _ := <- resChan
	//	if v == true {
	//		count++
	//	}
	//
	//	if count == 4 {
	//		close(primeChan)
	//		close(resChan)
	//		break
	//	}
	//}

	//for primeNumber := range primeChan {
	//	fmt.Println("primeNumber = ", primeNumber)
	//}

	for {
		primeNumber, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("prime = ", primeNumber)
	}

	fmt.Println("main end")

}