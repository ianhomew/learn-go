package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {

	// 假設客戶端訪問 轉 json 但是 客戶端希望 首字母 小寫開頭 怎麼辦
	// Ans: 使用「tag」
	// tag 可以透過反射機制獲取 常用場景是 序列化與反序列化

	var monster Monster = Monster{"牛魔王", 500, "芭蕉扇"}

	jsonByte, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("有錯")
	}

	jsonStr := string(jsonByte)
	fmt.Println(jsonStr)

}
