package main

import "fmt"

// go 內的繼承 是用 「匿名結構體」 在 struct 內嵌套一個匿名結構體
// 如上述 所以還有非匿名的結構體 這種模式叫做「組合」
// 匿名結構體可以直接訪問匿名結構體的屬性與方法 從而實現了繼承的特性

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods  // 匿名結構體在此 匿名結構體在此 匿名結構體在此 匿名結構體在此 匿名結構體在此
	Writer string
}

func main() {

	// 把 goods 寫在外面
	var goods = Goods{
		"我是一本書",
		300,
	}
	var book = Book{
		goods,
		"Egg",
	}

	fmt.Println("書名:", book.Name, "作者:", book.Writer)

	var book1 = Book{
		Goods{ // 把 goods 寫在裡面
			"秘密",
			99999,
		},
		"Secret",
	}
	//book1.Goods.Name 這樣也可以 可以清楚看到兩結構體的關係: book 內有 goods
	fmt.Println("書名:", book1.Name, "作者:", book1.Writer)

}
