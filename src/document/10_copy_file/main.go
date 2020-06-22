package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 自己編寫一個函數 接收兩個文件路徑
func CopyFile(srcFileName string, dstFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	// 通過 srcFile 取得 Reader
	reader := bufio.NewReader(srcFile)

	// 打開 dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
	}

	writer := bufio.NewWriter(dstFile)

	defer srcFile.Close()
	defer dstFile.Close()

	return io.Copy(writer, reader)

}

func main() {
	// 將一張圖片或電影或音樂檔案 複製到另一個文件
	// 使用 os.Copy()
	// 可以處理大文件
	// 因為有緩衝 一邊讀 一邊寫

	srcFile := "src/document/10_copy_file/resource/test.jpg"
	dstFile := "src/document/10_copy_file/des/"
	dstFile += "ddd" + ".jpg"

	_, err := CopyFile(srcFile, dstFile)
	if err == nil {
		fmt.Println("copy finished")
	} else {
		fmt.Println(err)
	}

}
