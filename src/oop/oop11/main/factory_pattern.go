package main

import (
	"fmt"
	"oop/oop11/model"
)

// go 沒有構造函數
// 當今天 有個 model 包 裡面有個 type student struct
// 因為是小寫所以是私有 引入了包也不能使用
// 怎麼辦？
// factory pattern 工廠模式！
func main() {
	//var s1 = model.Student{
	//	Name:  "Tom",
	//	Score: 44.5,
	//}
	//fmt.Println(s1)

	// 以上都沒問題

	// 接著把 model/student.go 內的 Student 改成 student

	var s1 = model.NewStudent("Claire", 60.66)
	fmt.Println(s1.Name) // 因為是結構體 可以直接用 s1
	fmt.Println((*s1).Name)

	fmt.Println(s1.Score)
	fmt.Println(s1.GetScore1()) // 寫一個方法 去讀取私有方法

}
