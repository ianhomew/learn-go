package main

import "fmt"

func main() {

	// var cities map[string]string    第一個[string]是 key 第二個 string 是 value

	// 1. 宣告 -> make -> 給值
	var cities map[string]string

	// 2. 宣告時就 make
	var cities1 = make(map[string]string)

	// 3. 宣告直接給值
	var cities2 = map[string]string{
		"c1": "台灣",
	}

	fmt.Println(cities)
	fmt.Println(cities1)
	fmt.Println(cities2)

	cities2["c2"] = "澳洲"
	fmt.Println(cities2["c1"])
	fmt.Println(cities2["c2"])
}
