package main

import "fmt"

// 先執行全局變數的定義
var n1 = test()

func test() int {
	fmt.Println("10") // first
	return 10
}

// 之後 main 調用前 會先調用init
func init() {
	fmt.Println("in init") // second
}

func main() {
	fmt.Println("in main") // third

	// 變數定義 >>> init >>> main
}
