package main

import (
	"fmt"
	"os"
)

func main() {

	var filePath string = "src/document/basic/data.txt"
	var file, err = os.Open(filePath)

	// file 的叫法
	// 1. file 對象
	// 2. file 指針
	// 3. file 文件句柄

	if err != nil {
		fmt.Println(err)
		return
	}

	// 輸出文件 結果發現 file 其實是個 pointer
	fmt.Printf("file = %v", file)

	err = file.Close()
	if err != nil {
		fmt.Println("close file error = ", err)
	}
}
