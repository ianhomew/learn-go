package main

import "fmt"

// slice of map: map 切片
// map 可以動態了

func main() {

	// 宣告一個 map 切片
	var monsters []map[string]string
	// 切片本身需要分配記憶體位址
	monsters = make([]map[string]string, 2)

	// 新增妖怪
	if monsters[0] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "妖精"
		monsters[0]["age"] = "500"
	}
	if monsters[1] == nil {
		monsters[0] = make(map[string]string)
		monsters[0]["name"] = "月娘"
		monsters[0]["age"] = "400"
	}

	fmt.Println(monsters)

}
