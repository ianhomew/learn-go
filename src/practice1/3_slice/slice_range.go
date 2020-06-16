package main

import "fmt"

// 如何對切片 slice 做 loop

func main() {

	var arr [6]int = [6]int{10, 20, 20, 30, 40, 50}

	slice := arr[1:4] // 20, 20, 30

	fmt.Println("===========================")

	// 1. use for
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}
	fmt.Println("===========================")

	// 2. use for range
	for index, value := range slice {
		fmt.Printf("slice[%v] = %v\n", index, value)
	}
	fmt.Println("===========================")

	// string
	var arr2 = [5]string{"grade A", "grade B", "grade C", "優良", "普通"} // array
	var strSlice []string = arr2[:]                                   // slice
	for i := 0; i < len(strSlice); i++ {
		fmt.Printf("strSlice[%v] = %v\n", i, strSlice[i])
	}
	fmt.Println("===========================")
	for index, value := range strSlice {
		fmt.Printf("strSlice[%v] = %v\n", index, value)
	}

	// 切片也可以再做切片
	strSlice2 := strSlice[2:len(strSlice)]
	fmt.Println(strSlice2)

	// 因為 slice 是引用類型 更改值的時候 全部的引用都會讀取到更改後的值
	strSlice2[0] = "改變改變改變了"
	fmt.Println(arr2)
	fmt.Println(strSlice)
	fmt.Println(strSlice2)

}
