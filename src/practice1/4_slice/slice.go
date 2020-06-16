package main

import "fmt"

// 切片 append 操作的本質就是對陣列的擴容

// 當切片大小不夠用的時候
// go 底層會建立一個新的陣列 newArr
// 接著將舊的陣列內的資料搬到 newArr
// 切片的指針重新指向新的陣列 newArr
// 以上我們都看不到

func main() {

	// append 函數 可以動態追加
	var intSlice = []int{1, 2, 3, 4, 5}

	var intSlice2 []int
	for i := 0; i < 5; i++ {
		intSlice2 = append(intSlice, i)
	}
	fmt.Println(intSlice2) // 123454  注意最後的結果

	for i := 0; i < 5; i++ {
		// 用原本的切片接收 才能追加
		intSlice = append(intSlice, i)
	}
	fmt.Println(intSlice) // 1234501234  注意最後的結果

	var intSlice3 []int
	if intSlice3 == nil {
		fmt.Println("現在的 intSlice3 == nil")
	}
	intSlice3 = append(intSlice, intSlice...)
	fmt.Println(intSlice3)

	// 切片擴容的證明
	var oldSlice = []int{100, 200}
	// oldSlice 的位址在 0xc000134140, len = 2, cap = 2
	fmt.Printf("oldSlice 的位址在 %p, len = %v, cap = %v\n", oldSlice, len(oldSlice), cap(oldSlice))
	// 丟一堆進去
	oldSlice = append(oldSlice, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	// oldSlice 的位址在 0xc00010e060, len = 11, cap = 12
	fmt.Printf("oldSlice 的位址在 %p, len = %v, cap = %v\n", oldSlice, len(oldSlice), cap(oldSlice))
}
