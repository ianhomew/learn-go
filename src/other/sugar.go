package main

import "fmt"

// Syntactic sugar 語法糖

// 注意 ... 的位置
// 函數使用: func way1(iamArr ...int)
// 呼叫時用: way2(sliceString...)

func way1(iamArr ...int) {
	for index, value := range iamArr {
		fmt.Println("index = ", index, "value = ", value)
	}
}

func way2(iamArr ...string) {
	for index, value := range iamArr {
		fmt.Println("index = ", index, "value = ", value)
	}
}

func main() {

	// 1. 參數的個數不確定的時候使用
	way1(1, 2, 3, 4, 5)

	// 2. 使用在切片 slice 上 可以將 slice 打散進行傳遞
	sliceString := []string{"apple", "boy", "cat", "dog"}
	way2(sliceString...) // 這裡跟函數反過來了

	sliceString1 := []string{"egg", "fox", "good", "high"}

	combineAll := append(sliceString, sliceString1...)
	fmt.Println(combineAll)

}
