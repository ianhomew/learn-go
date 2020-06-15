package main

import (
	"fmt"
	"unsafe"
)

// go 的陣列是 傳值 call by value 也就是丟進去函式內 會進行值拷貝

func main() {

	// 這樣定義後 chickenWeights 0~5 都有預設值
	var chickenWeights [6]float64 // 8 bytes

	chickenWeights[0] = 3
	chickenWeights[1] = 5
	chickenWeights[2] = 1
	chickenWeights[3] = 3.4
	chickenWeights[4] = 2
	chickenWeights[5] = 50

	sum := 0.0
	for i := 0; i < len(chickenWeights); i++ {
		sum += chickenWeights[i]
	}

	fmt.Println("total = ", sum)
	fmt.Printf("avg = %.2f\n", sum/float64(len(chickenWeights)))

	// 陣列第一個索引 跟 陣列本身開始的 地址 是相同的
	fmt.Printf("陣列本身開始的地址 = %p, 陣列第一個索引的地址 = %p\n", &chickenWeights, &chickenWeights[0])
	fmt.Printf("%p\n", &chickenWeights[1])

	// 利用強型別的陣列宣告 可以直接對記憶體位址做加總 從而得到其他陣列索引的值
	c := [3]int{1, 2, 3}                   // 宣告
	one := unsafe.Sizeof(c[0])             // 取得陣列 0 的大小 = 4 bytes
	d := uintptr(unsafe.Pointer(&c[0]))    // 取得陣列 0 的記憶體位址
	*(*int)(unsafe.Pointer(d + one)) = 999 // c[0]的記憶體位址加上 4 = c[1]
	fmt.Println("c =", c)

	// 陣列宣告方式
	var nums [2]int = [2]int{1, 2}
	fmt.Println(nums)
	var num1 = [2]int{1, 2}
	fmt.Println(num1)
	var num2 = [...]int{1, 2, 3, 4, 5} // ... 讓系統自己分配 是規定的寫法
	fmt.Println(num2)

	// 注意 陣列是從0 開始
	// 1:, 2:, 5:, 3: 是陣列的索引值
	var num3 = [...]int{1: 100, 2: 200, 5: 500, 3: 300}
	fmt.Println(num3)

}
