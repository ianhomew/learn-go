package main

import (
	"fmt"
	"strings"
)

// string 底層是 byte 陣列 所以 string 可以做切片
// string 是不可變的 所以不能 str[0] = 'A' 這樣寫
// 如果要修改字串 需要先將 string -> []byte or string -> []rune 然後再轉成字串

func main() {

	str := "Hello, world"
	slice := str[7:]
	fmt.Println("slice = ", slice)

	// 轉 byte
	arr1 := []byte(str)
	arr1[0] = 'Z'
	str = string(arr1)
	fmt.Println("str = ", str)
	// 上面如果有中文 會出現亂碼 可以使用 rune

	// 使用 rune
	str = "Hello, world"
	useRune := []rune(str)
	useRune[0] = '中'
	str = string(useRune)
	fmt.Println("str = ", str)

	str = "Hello, world"
	str = strings.Replace(str, "Hello", "Yo", 1)
	fmt.Println("str = ", str)

}
