package main

// 如果一個結構體內有個有名的結構體
// 這種模式做組合
// 如果是組合關係
// 那麼在訪問組合的結構體的屬性或方法時
// 必須帶上結構體名字

type A struct {
	Name string
}
type B struct {
	Name string
}
type C struct {
	a A // 有名字
	b B // 有名字
}

func main() {

	var c C
	c.a.Name = "a's name"
	c.b.Name = "b's name"
}
