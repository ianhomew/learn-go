package main

import (
	"fmt"
)

// close()
// close() 应当只由发送者执行，而不应由接收者执行
// 其效果是在最后发送的值被接收后停止该管道

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
