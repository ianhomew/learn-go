package utils

import "fmt"

func UtilTest() int {
	fmt.Println("in utils.test")
	return 0
}

func init() {
	fmt.Println("in utils.init")
}
