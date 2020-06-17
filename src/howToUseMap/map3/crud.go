package main

import "fmt"

func main() {

	myMap := make(map[string]string)

	// 新增
	myMap["n1"] = "N1"
	myMap["n2"] = "N2"
	fmt.Println("新增", myMap)

	// 修改
	myMap["n1"] = "Edited"
	fmt.Println("修改", myMap)

	// 刪除
	delete(myMap, "n2")

	// 當delete 的 key 不存在 則不操作 也不會報錯
	delete(myMap, "XXXXXXXXXXXXXX")
	fmt.Println("刪除", myMap)

	// 在 go 裡面如果要刪除所有的 key 只能
	// 1. 循環刪除
	// 2. make 一個新的
	//myMap = make(map[string]string)
	//fmt.Println(myMap) // result =  map[]

	// 查詢 找得到
	str, ok := myMap["n1"]
	if ok {
		fmt.Println("查詢 找得到", str)
	}
	// 查詢 找不到
	str1, ok1 := myMap["n999"]
	if !ok1 {
		fmt.Println("查詢 找不到", str1)
	}

}
