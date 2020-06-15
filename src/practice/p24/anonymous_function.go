package main

import "fmt"

var g = func(a, b int) int {
	return a + b
}

func main() {

	// 直接調用
	var s = func() string {
		return "直接調用"
	}()
	fmt.Println(s)

	// 可以多次調用
	var s1 = func() string {
		return "可以多次調用"
	}
	fmt.Println(s1())

	// 全局調用
	fmt.Println(g(1, 2))

}
