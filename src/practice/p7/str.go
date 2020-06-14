package main

import "fmt"

func main() {

	var str string = "hello"

	var c = str[0]

	fmt.Println(c)
	//str[0] = "a" // 錯 字串不能單獨修改 可以單獨配值

	// 使用反引號
	var strs string = `
A
B
C
\n A \n
`
	fmt.Printf(strs)

}
