package main

import (
	"fmt"
	"runtime"
)

func main() {

	// go 1.8 以後 不需要設置了
	// go 1.8 之前 需要設置 可以讓效率更好

	var cpuNum = runtime.NumCPU()
	fmt.Println("now cpu number = ", cpuNum)

	// 自己設置要用多少個 cpu
	runtime.GOMAXPROCS(cpuNum - 1)

	fmt.Println("finished")
}
