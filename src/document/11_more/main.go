package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount    int // 紀錄英文字數
	NumCount   int // 紀錄數字個數
	SpaceCount int // 紀錄空格
	OtherCount int // 紀錄其他
}

func main() {

	// 打開一個文件 建立一個 Reader
	// 每讀取一行 就去統計該行有多少個英文與數字空格等等
	// 然後保存到一個結構體內

	fileName := "src/document/11_more/abc.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var counts CharCount

	reader := bufio.NewReader(file)

	// 開始循環了
	for {
		str, err := reader.ReadString('\n')

		// str 不兼容中文
		//str1 := []rune(str) // 可以這樣

		for _, v := range str {
			//fmt.Println(v)
			switch { // 直接當作分支結構使用
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				counts.ChCount++
			case v == ' ' || v == '\t':
				counts.SpaceCount++
			case v >= '0' && v <= '9':
				counts.NumCount++
			case v == '\n':
				// 忽略掉
			default:
				counts.OtherCount++

			}
		}
		if err == io.EOF {
			fmt.Println("EOF", str)
			break
		}

	}

	fmt.Printf("ChCount = %d, NumCount = %d, SpaceCount = %d, OtherCount = %d",
		counts.ChCount, counts.NumCount, counts.SpaceCount, counts.OtherCount)

}
