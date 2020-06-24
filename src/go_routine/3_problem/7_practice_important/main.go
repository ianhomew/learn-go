package main

import (
	"fmt"
)

// close()
// close() 应当只由发送者执行，而不应由接收者执行
// 其效果是在最后发送的值被接收后停止该管道
// 管道的關閉 要在寫入的地方 寫完後 執行 這樣比較安全
// 因為在讀的地方關掉 不能保證 沒有協程在寫入了

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		intChan <- i
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
		fmt.Printf("read data = %v\n", v)
	}
	exitChan <- true
	// 全部讀完了就關閉 已經不用寫資料進去管道了
	close(exitChan)
}

func main() {

	// 寫資料的管道
	var intChan = make(chan int, 50)
	// 判斷是否讀完的管道
	var exitChan = make(chan bool, 1)

	go writeData(intChan)
	go readData(intChan, exitChan)

	// 如果下面都沒有程式 主線程就直接退出 兩個 go routine 就被銷毀了

	for {
		// 主線程幾乎是馬上執行到這邊 然後因為管道內無資料 就會進行等待(阻塞) 直到有資料進來
		// exitChan <- true  要等到有協程寫入資料  v, ok := <-exitChan 才能繼續往下執行
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
