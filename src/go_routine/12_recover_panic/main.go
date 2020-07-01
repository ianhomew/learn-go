package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello, world")
	}
}

func test() {
	// 處理錯誤 避免整個系統停止執行了
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 發生錯誤", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "ddd" // error here
}

func main() {

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}

}
