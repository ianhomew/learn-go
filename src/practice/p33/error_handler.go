package main

import "fmt"

// go 沒有 try catch finally
// go 用的是 defer, panic, recover
// go 中可以拋出一個 panic 的異常 然後在 defer 中通過 recover 捕獲這個異常 然後正常處理

func divideZero() {
	// 使用 defer + recover 處理異常
	defer func() {
		//err := recover() // 內建函數 可以捕獲到異常

		// 初始化 err         判斷是否 nil
		if err := recover(); err != nil { // 表示有錯誤
			// 可以做處理 例如發送mail給管理員等等
			fmt.Println("err = ", err)
		}

	}()

	n1 := 10
	n2 := 0
	divide := n1 / n2

	fmt.Println("divide = ", divide)

}

func main() {

	divideZero()

	// 沒捕貨做處理 就不會跑這邊了
	fmt.Println("after divideZero()")
}
