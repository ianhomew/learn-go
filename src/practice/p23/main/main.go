package main

import (
	"fmt"
	"practice/p23/utils"
)

// 會發現執行順序是 先載入 "practice/p23/utils" 裡面有 init
// 所以會先執行 utils.init
// 接著 執行 var s = utils.UtilTest()
// 接著執行 main.init
// 最後執行 main

//var s = test()
var s = utils.UtilTest()

func init() {
	fmt.Println("in main.init")
}

func test() string {
	fmt.Println("in main.test")
	return ""
}

func main() {

	fmt.Println("in main")

}
