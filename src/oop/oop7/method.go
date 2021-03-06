package main

import "fmt"

// golang 中 方法是作用在指定的資料型別上(也就是 和指定的資料型別綁定)

type Person struct {
	Name string
}

// 只有 type A 可以用 test()
// p 是一個拷貝
// 這樣是函數
func (p Person) test() {
	p.Name = "Changed name"
	fmt.Println("test()  = ", p.Name)
}

// 這樣是方法
func test() {
	fmt.Println("跟上面比較差異 -> (a, A)")
}

func main() {

	var p Person
	p.Name = "Stella"

	fmt.Println("main name = ", p.Name)
	p.test()
	fmt.Println("main name = ", p.Name)
}
