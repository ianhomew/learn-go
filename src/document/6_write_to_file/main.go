package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打開一個存在的文件 在原本的內容上追加內容: "I AM APPEND"

	filePath := "src/document/6_write_to_file/abc.txt"
	// 打開已經存在的文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// write
	str := "I AM APPEND"
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
