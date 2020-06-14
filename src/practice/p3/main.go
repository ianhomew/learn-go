package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {

	// 直接設定好類型大小
	var i int8 = 10
	var j int32 = 10 // 強制int32
	var k int64 = math.MaxInt64

	fmt.Printf("type of i = %T, 字節大小 = %d \n", i, unsafe.Sizeof(i))
	fmt.Printf("type of j = %T, 字節大小 = %d \n", j, unsafe.Sizeof(j))
	fmt.Printf("type of k = %T, 字節大小 = %d \n", k, unsafe.Sizeof(k))

	// 由系統設定類型
	var x = 10
	var y = 10 // 預設給 int64
	var z = math.MaxInt64

	fmt.Printf("type of x = %T, 字節大小 = %d \n", x, unsafe.Sizeof(x))
	fmt.Printf("type of y = %T, 字節大小 = %d \n", y, unsafe.Sizeof(y))
	fmt.Printf("type of z = %T, 字節大小 = %d \n", z, unsafe.Sizeof(z))

}
