package main

import "fmt"

// 結構體是用戶自訂義的資料型別 和其他資料型別進行轉換的時候
// 需要有「完全相同」的屬性(名字 個數 類型)

type A struct {
	Num int
}
type B struct {
	Num int
}

func main() {

	var a A
	var b B

	// 強轉
	a = A(b) // 這樣可以

	fmt.Println(a, b)

}
