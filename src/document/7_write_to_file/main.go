package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打開一個存在的文件 將原來的內容讀出來顯示在終端 並且追加五句 "I love MMM"

	filePath := "src/document/7_write_to_file/abc.txt"
	// 打開已經存在的文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// read and print
	reader := bufio.NewReader(file)
	for {
		readString, err := reader.ReadString('\n') // 每讀到一個 \n 取出來
		if err == nil && err != io.EOF {
			fmt.Print(readString)
		} else {
			break
		}
	}

	// write
	str := "I love MMM"
	//
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(str + "\r\n")
	}

	// 因為 writer 是帶緩存的 因此在調用 WriteString 方法時 事先寫入緩存的
	// 看結構體
	//type Writer struct {
	//	err error
	//	buf []byte // 都是先寫在這邊的
	//	n   int
	//	wr  io.Writer
	//}
	err = writer.Flush()

}
