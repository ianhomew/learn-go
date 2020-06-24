package main

import (
	"fmt"
	"time"
)

func main() {

	// 開啟一個協程 每 1 秒輸出 hello world
	// 主線程 每秒輸出 hello golang 輸出 10 次後退出

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("hello world")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("hello golang")
		time.Sleep(time.Second)
	}

}
