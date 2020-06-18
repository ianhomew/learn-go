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
		monsters[1] = make(map[string]string)
		monsters[1]["name"] = "月娘"
		monsters[1]["age"] = "400"
	}

	fmt.Println(monsters)

	// 以上程式碼都是正確無誤的 但是有個問題 如下:
	// 會報 panic: runtime error: index out of range [2] with length 2
	//if monsters[2] == nil {
	//	monsters[2] = make(map[string]string)
	//	monsters[2]["name"] = "月娘"
	//	monsters[2]["age"] = "400"
	//}

	// append 出現！ 使用 slice 的 append 函數
	// 1. 先定義一個 monster
	newMonster := map[string]string{
		"name": "新妖怪",
		"age":  "200",
	}
	// 2. append
	monsters = append(monsters, newMonster)
	fmt.Println(monsters)

	// 可不可以都用 append 不管 make 時分配的大小？ Ans: 可以
	var pokemons []map[string]string
	pokemons = make([]map[string]string, 1)

	//pokemon[0] = make(map[string]string)
	//pokemon[0]["name"] = "皮卡丘"
	//pokemon[0]["age"] = "5"
	pokemons = append(pokemons, map[string]string{
		"name": "皮卡丘",
		"age":  "5",
	})
	pokemons = append(pokemons, map[string]string{
		"name": "雷丘",
		"age":  "15",
	})

	fmt.Println(pokemons)

}
