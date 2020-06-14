package main

import "fmt"

// args is a slice
// 支持0 到 N個參數數
func sum(args ...int) (sum int) {
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return
}

// sum1 至少需要一個參數 最多 N個參數數
func sum1(n1 int, args ...int) (sum int) {
	sum = n1
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return
}

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum1(1, 2, 3, 4, 5))
}
