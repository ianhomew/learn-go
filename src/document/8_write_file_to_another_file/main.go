package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 將一個文件的內容 寫入到另外一個文件 這兩個文件已經存在

	// 1. 先讀取 original.txt 資料讀取到內存
	// 2. 將讀取到的內容寫入 des.txt

	originalFile := "src/document/8_write_file_to_another_file/original.txt"
	desFile := "src/document/8_write_file_to_another_file/des.txt"

	data, err := ioutil.ReadFile(originalFile)
	if err != nil {
		fmt.Println("has error")
		return
	}

	err = ioutil.WriteFile(desFile, data, 0666)
	if err != nil {
		fmt.Println("has error")
		return
	}

}
