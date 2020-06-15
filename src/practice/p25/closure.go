package main

import "fmt"

// Add() 是一個函數 返回的資料類型是 func (int) int
func Add() func(int) int {
	n := 10 // 定義一個變數n 給值 10
	// 最後返回一個函數
	return func(i int) int {
		n += i
		return n
	}
}

// 定義一個函數名稱叫做 Sub() 返回兩個函數  func(int) int       func() int
// 兩個以上的參數需要用括弧刮起來
// go閉包跟js閉包比較像 在這裡定義了初始變數 n 為 100, programmer 不希望 n 被隨便賦值 只希望開放兩個功能
// 遞減的功能 與 重新設定n 的功能
func Sub() (func(int) int, func() int) {
	n := 100
	return func(i int) int {
			n -= i
			return n
		}, func() int {
			n = 100
			return n
		}
}

func main() {

	var s = func(a, b int) string {
		sum := a + b
		return func(a, b, sum int) string {
			return fmt.Sprintf("%v + %v = %v", a, b, sum)
		}(a, b, sum)
	}
	fmt.Println(s(1, 2))

	// add 是一個函數
	var add = Add()
	fmt.Println(add(1))
	fmt.Println(add(1))

	var sub, reset = Sub()
	fmt.Println(sub(1))  // 99
	fmt.Println(sub(1))  // 98
	fmt.Println(reset()) // 100
	fmt.Println(sub(1))  // 99
	fmt.Println(sub(1))  // 98

}
