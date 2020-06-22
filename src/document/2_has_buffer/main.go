package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 帶緩衝區的方式
	// 更加適合讀取大文件
	// os.Open, fuke.Close,

	var filePath string = "src/document/has_buffer/data.txt"
	var file, err = os.Open(filePath)

	// file 的叫法
	// 1. file 對象
	// 2. file 指針
	// 3. file 文件句柄

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// 建立一個 *Reader 是帶緩衝的 默認 4096 bytes
	// 有緩出的好處: 不是一次讀取整個文件
	reader := bufio.NewReader(file)
	for {
		// The best way to read a file line by line is to use reader.ReadString('\n')
		str, err := reader.ReadString('\n') // 表示一行一行讀 \n = 換行
		if err == nil {
			if err == io.EOF { // EOF = end of line 整個文件讀完了
				break
			}
		}

		// output
		fmt.Print(str)
	}

}
