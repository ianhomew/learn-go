package main

import "fmt"

// 結構體內如果有兩個以上的匿名結構體
// 且這些結構體有相同的屬性與方法(同時結構體本身沒有)
// 在訪問的時候就必須明確指定匿名結構體的名字
// 否則編譯會報錯

type A struct {
	Name string
}
type B struct {
	Name string
}

// 如果一個結構體裡面有多個匿名結構體
// 那麼該結構體可以直接訪問裡面的匿名結構體的屬性與方法
// 從而實現了 「多重繼承」    盡量不要使用多重繼承 會變得很複雜
type C struct {
	A
	B
}

func main() {

	var c C
	// 直接報錯 編譯器不知道你是要哪個 Name
	//c.Name = "Bob"

	// 這樣就好嚕
	c.A.Name = "A's name"
	c.B.Name = "B's name"

	// 但是如果 C 結構體內本身含有 Name string
	// 這時候使用 c.Name 就不會報錯

	fmt.Println(c)
}
