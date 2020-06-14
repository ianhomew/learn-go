package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	// 基本資料類型轉成字串
	// 1. 使用 fmt.Sprintf
	// 2. 使用 strconv 包

	var num1 int = 10
	var num2 float64 = 1.0001
	var b bool = false
	var by byte = 'a'
	var str string

	// 第一種方式
	fmt.Println("第一種方式")

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str=%q\n", str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str=%q\n", str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str=%q\n", str)

	// 無法輸出 a
	str = fmt.Sprintf("%v", by) // %v	值的默认格式表示
	fmt.Printf("str=%q\n", str)

	str = fmt.Sprintf("%c", by) // %c	该值对应的unicode码值
	fmt.Printf("str=%q\n", str)

	// 第二種方式
	fmt.Println("第二種方式")

	fmt.Println(strconv.FormatInt(int64(num1), 10))
	fmt.Println(strconv.FormatFloat(num2, 'f', 10, 64))
	fmt.Println(strconv.FormatFloat(num2, 'f', -1, 64))
	fmt.Println(strconv.FormatBool(b))

	fmt.Println("OTHER")

	// strconv.Itoa
	var i int = 456789
	fmt.Println(strconv.Itoa(i))
	var j int64 = math.MaxInt64
	fmt.Println(j)
	fmt.Println(strconv.Itoa(int(j)))

}
