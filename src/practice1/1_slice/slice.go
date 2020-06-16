package main

import "fmt"

// 一定要理解 slice 的底層運作原理
// slice 是引用類型

func main() {

	// 定義陣列
	var intArr [5]int = [...]int{1, 22, 33, 44, 55}

	// 定義 slice
	// 1 表示從索引為 1 開始到 索引為 3 的結束且不包含 3
	// intArr[1:3] 解釋為 slice 引用到 intArr 陣列裡面的 intArr[1], intArr[2]
	slice := intArr[1:3]
	fmt.Println("intArr = ", intArr, "len = ", len(intArr))
	fmt.Println("slice = ", slice, "len = ", len(slice)) // 22, 33
	fmt.Println("slice 的容量 = ", cap(slice))              // 切片是可以動態變化的 cap = capability

	var s []int = []int{1, 2, 3, 4, 5}

	fmt.Printf("type = %T\n", s)
	fmt.Println("容量 = ", cap(s))

	// 以下非常重要
	// slice 是引用類型 所以存的是 「被引用的類型的位址」
	slice[0] = 9999 // 引用都會改
	fmt.Println("========================")
	fmt.Printf("slice[0]位址 = %p, 值 = %v\n", &slice[0], slice[0])
	fmt.Printf("intArr[1]位址 = %p, 值 = %v\n", &intArr[1], intArr[1])
	// &slice[0] 與 &intArr[1] 記憶體位址相同
	fmt.Println("========================")

	//s1 := intArr[:]  // 全部
	//s2 := intArr[1:] // 等於 intArr[1:len(intArr)]
	//s3 := intArr[:2] // 等於 intArr[0:2]

	// new 用來分配「值類型」 int, float32, struct...
	// make 用來分配「引用類型」 chan, map slice...

	// 所以可以這樣用
	//useMake := make([]int, 0, 5)

	// 切片 其實就是一個結構體 長這樣
	//type slice struct {
	//	ptr *[2]int // *[2]int 會根據不同切片類型而變 ptr 會指向真正存放資料的記憶體位置 該位址為陣列
	//	len int // 記錄空間大小
	//	cap int // 紀錄容量
	//}

}
