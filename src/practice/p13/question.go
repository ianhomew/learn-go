package main

import "fmt"

func main() {

	// 不用中間變數將兩個變數做交換
	a := 10
	b := 20

	// 開始
	// 第一種
	//fmt.Println("Old : ", a, b)
	//a, b = b, a
	//fmt.Println("New : ", a, b)

	// 第二種
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("New : ", a, b)

}
