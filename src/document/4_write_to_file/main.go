package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	filePath := "src/document/4_write_to_file/abc.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	// write
	str := "Hello, Gardon!"
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
