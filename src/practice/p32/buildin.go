package main

import "fmt"

// 內建函數
func main() {

	// len 跳過~

	// new 用來分配「值類型」 int, float32, struct...
	n1 := 100
	// 類型=int, 值=100, 位址=0xc000018070
	fmt.Printf("n1   類型=%T, 值=%v, 位址=%v\n", n1, n1, &n1)

	n2 := new(int) // *int
	// n2   類型=*int, 值=0xc000018088, 位址=0xc00000e030, 指向的值=0
	fmt.Printf("n2   類型=%T, 值=%v, 位址=%v, 指向的值=%v\n", n2, n2, &n2, *n2)

	// make 用來分配「引用類型」 chan, map slice...
	// 目前還不會 先跳過

}
