package main

import (
	"fmt"
)

// 寫一個function 印出參數類型 並且印出值

func PrintType(data ...interface{}) {
	for _, item := range data {
		switch v := item.(type) { // type 是關鍵字 固定的寫法
		case int:
			fmt.Printf("type = %T, value = %v\n", v, v)
		case bool:
			fmt.Printf("type = %T, value = %v\n", v, v)
		case float64:
			fmt.Printf("type = %T, value = %v\n", v, v)
		case nil:
			fmt.Printf("type = %T, value = %v\n", v, v)
		case string:
			fmt.Printf("type = %T, value = %v\n", v, v)
		default:
			fmt.Printf("type = %T, value = %v\n", v, v)
		}
	}
}

func main() {

	PrintType(100)
	PrintType(true)
	PrintType(10.565)
	PrintType(nil)
	PrintType("台灣")

}
