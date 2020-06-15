package main

import "fmt"

var s = test()

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
