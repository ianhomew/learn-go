package main

import "fmt"

// 切片的使用 有三種方式

func main() {

	// 1. 定義一個切片 讓切片去引用一個已經存在的陣列
	var arr = [...]int{1, 2, 3} // arr 是陣列
	var way1 []int = arr[:]     // way1 是結構體
	fmt.Println("第一種方式", way1)

	fmt.Println("=======================")
	// 2. 使用 make
	// 切片必須 make 後才能使用
	var slice []float64 // 只宣告變數為一個 float64 的切片 這時候還不能使用
	//slice[0] = 100 // 會錯 因為 slice == nil
	fmt.Printf("只有宣告變數 還沒分配記憶體位址 silce == nill ?  ans: %t\n", slice == nil)
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))

	fmt.Println("接著分配記憶體空間")
	slice = make([]float64, 5, 10)
	fmt.Println("slice = ", slice, "len = ", len(slice), "cap = ", cap(slice))

	// 面試常考: 第一種與第二種的差異在於
	// 方式1 直接引用陣列 這個陣列是事先存在的 工程師都可以看到
	// 方式2 是透過 make 建立切片 而 make 也會建立一個陣列 該陣列由切片做維護 我們看不到

	// 3. 定義切片的時候直接指定具體的陣列 使用原理類似 make
	var way3 []int = []int{1, 2, 3}
	fmt.Println("way_3 = ", way3)

}
