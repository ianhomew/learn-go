package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var str = "this is string"
	fmt.Println(str)

	var a, b int32 = 10, 20
	fmt.Println(a, b)

	var uninit int
	fmt.Println(uninit)

	// := 是聲明語句 宣告的意思 var簡寫
	//var t float32 = 13.4
	//fmt.Println(t)
	t := 10.3
	fmt.Println(t)

	// 這樣可以過 因為有宣告新的變數
	// 沒有宣告格式 只能在function內使用
	t, tt := 10.4, 10.5
	fmt.Println(t, tt)

	n := t
	fmt.Println(n)

	const constValue = "CONST"
	fmt.Println(constValue)

	const (
		A int32 = 10
		B       = 20
		C       = "TEST"
		D       = len(C)
		E       = unsafe.Sizeof(C)
	)
	fmt.Println(A, B, C, D, E)

	const (
		aa = iota // iota 第一次使用是0
		bb
		cc
	)
	fmt.Println(aa, bb, cc)

	ccc := 1
	//ddd := 5
	//ccc = ddd++
	fmt.Println(ccc)

	fmt.Println("GET MAX", max(10, 20))

	var str1 string = "HE IS"
	var str2 string = "GOD"

	var str_1, str_2 = swap(str1, str2)

	fmt.Println("SWAP", str_1, str_2)

}

func max(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}

func swap(str1 string, str2 string) (string, string) {
	return str2, str1
}
