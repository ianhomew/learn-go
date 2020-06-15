package main

import (
	"fmt"
	"time"
)

// 展示日期與時間相關的函數方法
func main() {

	// 1. get Now
	now := time.Now() // return a type 'Time'
	fmt.Printf("%T\n%v\n", now, now)

	fmt.Println("年", now.Year())
	fmt.Println(now.Month())
	fmt.Println(int(now.Month()))
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println("====================")

	// 2. 格式化日期時間
	// 也可以用 Sprintf
	// %02d 表示輸出整數 兩位數 不滿則捕0
	var dateString = fmt.Sprintf("當前: %02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println("格式化後 ", dateString)

	fmt.Println("另一種格式化方式")
	// 發明語言 就是可以任性
	// 2006-01-02 15:04:05 這個字串是固定的 必須這樣寫.................WTF
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now.Format("15:04:05"))

}
