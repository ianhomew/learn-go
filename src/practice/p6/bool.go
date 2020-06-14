package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// bool 佔 1 個字節 (byte)
	var b = true
	fmt.Printf("type of b = %T, 字節大小 = %d \n", b, unsafe.Sizeof(b))

}
