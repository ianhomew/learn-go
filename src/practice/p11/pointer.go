package main

import "fmt"

func main() {

	var i int = 10
	fmt.Println("i 的記憶體位址是 ", &i)

	// ptr 是一個指針變數
	// ptr 的變數類型是 *int
	// ptr 本身的值是 &i
	var ptr *int = &i

	i++
	fmt.Println("ptr 的記憶體位址是 ", &ptr)
	fmt.Println(ptr, *ptr)

	// 值類型 通常儲存在棧區
	// 引用類型 通常儲存在堆區
	// 「通常」不是絕對 因為有go 編譯有「逃逸分析」

}
