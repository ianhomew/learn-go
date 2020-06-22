package main

import (
	"fmt"
	"os"
)

// os.Open
// func Open(name string) (file *File, err error)
// Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；
// 对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。

// 注意
// 最後需要將 file.Close()

func main() {

	var filePath string = "src/document/basic/data.txt"
	var file, err = os.Open(filePath)

	// file 的叫法
	// 1. file 稱作對象
	// 2. file 稱作指針
	// 3. file 稱作文件句柄

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
