package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func sum1(cFun func(int, int) int, n1 int, n2 int) int {
	return cFun(n1, n2)
}

func main() {

	f := sum
	f1 := sum

	res := f(1, 2)

	fmt.Println(res)
	if &f == &f1 {
		fmt.Println("f == f1")
	} else {
		fmt.Println("f != f1")
	}

	fmt.Printf("f type = %T \n", f)

	// 傳函數
	f2 := sum
	fmt.Println(sum1(f2, 10, 20))

	// go 可以自定義型別
	type myInt int
	var cN1 myInt
	cN1 = 10
	fmt.Println("cN1", cN1)

}
