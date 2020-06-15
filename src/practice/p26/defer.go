package main

import "fmt"

func sum(a, b int) int {
	// 面試常問
	// 棧 = stack
	// 編譯器看到 defer 會把 defer 後面的語句 放入獨立的棧區(專門放 defer 的) 不會馬上執行 並且會將相關會用到的值複製一份過去(call by value)
	// 當函數執行完畢後 再從 defer 棧中 依照先入後出的方式出棧 並執行之
	defer fmt.Println("last")
	defer func(a int) {
		fmt.Println("a = ", a)
	}(a)

	a++
	a++
	fmt.Println("before return a = ", a)

	defer fmt.Println("============")
	return a + b
}

func main() {

	a := 10
	b := 20

	sum := sum(a, b)

	fmt.Println("sum = ", sum)
}
