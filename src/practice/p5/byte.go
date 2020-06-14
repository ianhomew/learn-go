package main

import (
	"fmt"
)

func main() {

	var s1 byte = 'a'
	var s2 byte = '0'

	// out put is ASCII code
	fmt.Println("s1=", s1)
	fmt.Println("s2=", s2)

	// 輸出原樣 使用 %c
	// $c = the character represented by the corresponding Unicode code point
	fmt.Printf("s1=%c\n", s1)
	fmt.Printf("s2=%c\n", s2)

	//var s3 byte = '北' // overflow
	//fmt.Printf("s3=%d\n", s3)

	// 字符常量 是用單引號刮起來的
	var s4 int = '北' // golang 編碼採用 utf8     utf8 編碼是 ascii 的擴充 所以可以兼容
	fmt.Printf("s4=%c   s4=%d\n", s4, s4)
	var s5 int = 'a' // 97 -> 對應 ascii 與 utf8 的值
	var s6 = s5 + 1  // 98 -> b
	fmt.Printf("s6 = %d\n", s6)
	fmt.Printf("s6 = %c\n", s6)

}
