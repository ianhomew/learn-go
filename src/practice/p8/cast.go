package main

import "fmt"

func main() {

	// golang 不會自動轉換數據類型
	var i int = 8
	var j = int64(i)

	fmt.Println(j)

}
