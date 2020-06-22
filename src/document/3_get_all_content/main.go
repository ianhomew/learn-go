package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// 一次性的讀取整個檔案 ioutil.ReadFile
	// 檔案太大 記憶體會爆掉
	// ioutil.ReadFile 不用我們處理 Open/Close 包已經幫我們處理好

	var filePath string = "src/document/3_get_all_content/data.txt"
	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%v", file)
	fmt.Printf("%s", file)
}
