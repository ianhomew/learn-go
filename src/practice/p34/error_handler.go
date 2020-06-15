package main

import (
	"errors"
	"fmt"
)

// 也支持自定義錯誤 使用 errors.New 和 panic 內建函數
// panic 接收 interface{} 就是任何值！！！ 可以接收 error 類型的變數 輸出錯誤訊息後會退出程序

// 讀取一個配置文件 init.conf
// 有錯誤則返回自定義錯誤

func readConf(name string) (err error) {

	if name == "config.ini" {
		// 讀取
		return nil
	} else {
		err := errors.New("文件名錯誤")
		return err
	}
}

func good() {
	err := readConf("config.ini")
	if err != nil {
		// 錯誤則輸出並終止
		panic(err)
	}

	fmt.Println("test1 後面的程式碼")
}
func bad() {
	err := readConf("config11111111111.ini")
	if err != nil {
		// 讀取錯誤 輸出並終止
		panic(err)
	}

	fmt.Println("test1 後面的程式碼")
}

func main() {

	good()
	bad()
	fmt.Println("main 後面的程式碼")
}
