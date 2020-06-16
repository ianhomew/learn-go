package main

import "fmt"

// slice 是引用類型
func testSlice(slice []int) {
	slice[len(slice)-1] = 3333
}

func main() {

	// slice 是 引用類型

	var intArr = []int{100, 200, 300, 400, 500}
	var slice []int

	slice = intArr[:]

	var slice1 = slice

	slice[0] = 9999

	fmt.Println("intArr = ", intArr)
	fmt.Println("slice = ", slice)
	fmt.Println("slice1 = ", slice1)
	//intArr =  [9999 200 300 400 500]
	//slice =  [9999 200 300 400 500]
	//slice1 =  [9999 200 300 400 500]

	fmt.Println("=====================")

	testSlice(slice)
	fmt.Println("intArr = ", intArr)
	fmt.Println("slice = ", slice)
	fmt.Println("slice1 = ", slice1)

}
