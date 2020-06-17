package main

import "fmt"

// 二維的 map

func main() {

	studentMap := make(map[string]map[string]string)

	// 這裡還要 make 一次 沒有 make 會報 panic 錯誤
	studentMap["s_001"] = make(map[string]string, 2)
	studentMap["s_001"]["name"] = "Tom"
	studentMap["s_001"]["sex"] = "Boy"

	studentMap["s_002"] = make(map[string]string, 2)
	studentMap["s_002"]["name"] = "Marry"
	studentMap["s_002"]["sex"] = "Girl"

	fmt.Println(studentMap)

	fmt.Println("studentMap len = ", len(studentMap))
}
