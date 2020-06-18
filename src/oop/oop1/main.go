package main

import "fmt"

// 引用類型的資料型別 都需要 make ！！！！！！！！！

// slice, map, pointer 默認值是 nil 表示尚未分配記憶體空間

// 定義一個結構體 結構體是值類型 因為是直接指向 記憶體空間 不是放記憶體位址
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
	Other [3]int
}

func main() {

	// 結構體 -> 資料的管理與維護 (本身是一個整體)
	// "go 支持物件導向的「特性」"
	// go 是基於結構體 struct 來實現 oop

	// go 是接口導向的 「面向接口」 OOP 大陸稱為「面向對象」

	// 結構體 就像 class  可以建立多個物件

	// 以下開始 以上廢話
	var cat1 Cat
	cat1.Name = "灰色的貓"
	cat1.Age = 20
	cat1.Color = "灰色的"
	fmt.Println(cat1)
	fmt.Printf("%p\n", &cat1)

	// 不同結構體的變數是獨立的 因為結構體是「值類型」
	var cat2 Cat
	cat2.Name = "白色的貓"
	cat2.Age = 3
	cat2.Color = "白色的"
	fmt.Println(cat2)
	fmt.Printf("%p\n", &cat2)

	var cat3 Cat = cat2 // 因為是值類型 所以是整個複製一份 所以是另外分配了一個記憶體空間
	fmt.Println(cat3)
}
