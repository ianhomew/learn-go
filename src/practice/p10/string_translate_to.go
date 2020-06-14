package main

import (
	"fmt"
	"strconv"
)

func main() {

	var str = "false"
	var b bool

	b, _ = strconv.ParseBool(str)
	fmt.Printf("type = %T, value = %v\n", b, b)

	var s1 = "123456"
	var n1, _ = strconv.ParseInt(s1, 10, 64)

	fmt.Printf("type = %T, value = %v\n", s1, n1)

	var s2 = "10.0000015647897"
	var f1, _ = strconv.ParseFloat(s2, 64)
	fmt.Printf("type = %T, value = %v\n", f1, f1)

	// 遇到不能轉成數字 會傳回默認值
	var e = "hello"
	var n2, err = strconv.ParseInt(e, 10, 64)
	fmt.Printf("type = %T, value = %v\n", n2, n2)
	fmt.Println(err) // 會有錯誤的

	// 會傳回默認值
	var b1, err1 = strconv.ParseBool(e)
	fmt.Printf("type = %T, value = %v\n", b1, b1)
	fmt.Println(err1) // 會有錯誤的

}
