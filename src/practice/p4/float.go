package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// 精度損失的問題

	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	var num3 = 1.1
	fmt.Println("num1 = ", num1)
	fmt.Println("num2 = ", num2)

	// golang 預設給 float64
	fmt.Printf("type of x = %T, 字節大小 = %d \n", num3, unsafe.Sizeof(num3))

}
