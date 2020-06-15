package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

// 定義一個型別為 `func(int, int) int`
type myCustomFun func(int, int) int

func sum2(hana myCustomFun, n1 int, n2 int) int {
	return hana(n1, n2)
}
func main() {

	// 在外面定義一個類型 type myCustomFun 的函數
	fmt.Println(sum2(sum, 100, 200))
}
