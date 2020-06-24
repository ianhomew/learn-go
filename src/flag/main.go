package main

import (
	"flag"
	"fmt"
)

// go build -o bin/main bin/main src/flag/main.go

func main() {

	// 定義變數 用於接收命令行的參數值
	var user, pwd, host string
	var port int

	// 連接 mysql
	// -u root -pwd root -h 127.0.0.1 -port 3306

	flag.StringVar(&user, "u", "", "使用者默認為空")
	flag.StringVar(&pwd, "pwd", "", "密碼")
	flag.StringVar(&host, "h", "localhost", "主機")
	flag.IntVar(&port, "port", 3306, "port")

	// 必須調用此方法！！！！！！！！！！！
	flag.Parse()

	fmt.Printf("user = %v, pwd = %v, host = %v, port = %v",
		user, pwd, host, port)

}
