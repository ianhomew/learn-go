package main

import (
	"fmt"
)

// 如果只是向管道寫入數據 「沒有讀取」
// 就會出現阻塞而 dead lock
// 原因是 intChan 容量是 10 而 writeData 會寫入 50個數據
// 寫跟讀管道的速度不同 不會出現 dead lock 錯誤

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i // 寫50 個 但是容量只有 10 又沒有從管道取資料出來 會阻塞
		fmt.Println("write data =", i)
	}
	// 全部寫完了就關閉 已經不用寫資料進去管道了
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		// 有讀出來 就不會有問題 就算讀得慢
		//time.Sleep(5 * time.Second)
		fmt.Printf("read data = %v\n", v)
	}
	exitChan <- true
	// 全部讀完了就關閉 已經不用寫資料進去管道了
	close(exitChan)
}

func main() {

	// 寫資料的管道
	var intChan = make(chan int, 10)
	// 判斷是否讀完的管道
	var exitChan = make(chan bool, 1)

	go writeData(intChan)
	// 完全沒有讀 會錯
	//go readData(intChan, exitChan)
	go readData(intChan, exitChan)

	// 上面 只有寫入管道 而沒有讀取管道內的數據 滿了就爆了
	// fatal error: all goroutines are asleep - deadlock!

	// 如果下面都沒有程式 主線程就直接退出 兩個 go routine 就被銷毀了

	for {
		v, ok := <-exitChan
		fmt.Println(v, ok)

		// 這是根據寫入管道的值判斷
		if v == true {
			break
		}

		// 根據管道內 沒有資料 的判斷
		//if !ok {
		//	break
		//}
	}

	fmt.Println("安全結束")

}
