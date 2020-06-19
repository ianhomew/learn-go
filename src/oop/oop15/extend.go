package main

import "fmt"

type Goods struct {
	Name  string
	Price int
}

type Book struct {
	Goods  // 匿名結構體在此 匿名結構體在此 匿名結構體在此 匿名結構體在此 匿名結構體在此
	Writer string
}

func (g *Goods) PrintName() {
	fmt.Println("Goods PrintName", g.Name)
}

// 這個可以註解掉 book1.PrintName() 會有不同結果
func (b *Book) PrintName() {
	fmt.Println("Book PrintName", b.Name)
}

func main() {

	var book1 = Book{
		Goods{ // 把 goods 寫在裡面
			"秘密",
			99999,
		},
		"Secret",
	}

	// 因為 book 類型是 Book 所以會先在 Book 結構體內尋找 PrintName()
	// 如果 Book 沒有 PrintName() 就會去匿名結構體內找是否有 PrintName()
	// 都找不到就報錯嚕
	// 「就近訪問原則」
	book1.PrintName()
	// 直接指定訪問  Goods 內的 PrintName() 方法
	book1.Goods.PrintName()

}
