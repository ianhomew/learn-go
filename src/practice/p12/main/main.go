package main

import (
	"fmt"

	// 引入包 就是指定到資料夾的目錄下(other) 在該直屬目錄下的 .go 檔案內的 大寫開頭的變數或方法都可以使用
	// 在資料夾(other)底下的資料夾(dir2)裡面的檔案是無法引用的 必須額外 import
	"practice/p12/other"
	"practice/p12/other/dir2"
	"strconv"
)

func main() {
	fmt.Println(other.HeroName)

	i := other.Add(4, 4)
	fmt.Println(strconv.FormatInt(int64(i), 10))

	j := dir2.Sub(10, 4)
	fmt.Println(strconv.FormatInt(int64(j), 10))

}
