package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打開一個存在的文件 將原本的內容覆蓋掉 並且寫 10 句 Hello, world!

	filePath := "src/document/5_write_to_file/abc.txt"
	// 打開已經存在的文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// write
	str := "Hello, world!"
	//
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
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
