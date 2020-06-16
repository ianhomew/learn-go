package main

import "fmt"

func main() {

	var slice = []int{100, 200, 300, 400, 500}

	var newSlice = make([]int, 1)
	// newSlice =  [0]
	fmt.Println("newSlice = ", newSlice)

	copy(newSlice, slice)

	// newSlice =  [100]
	fmt.Println("newSlice = ", newSlice)

}
