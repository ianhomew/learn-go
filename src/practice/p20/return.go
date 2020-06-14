package main

import "fmt"

// cal function has two params, a and b, and has two return values named sum and sub
// sum and sub are already declared
func cal(a int, b int) (sum int, sub int) {

	// sum := a + b // 錯誤 因為宣告函數的時候已經宣告好了
	sum = a + b
	sub = a - b

	return
}

func main() {

	sum, sub := cal(100, 20)

	fmt.Println("sum = ", sum, "sub = ", sub)

}
