package main

import (
	"fmt"
	"strconv"
)

func main() {

	var s = "hello世界"
	fmt.Println("s len = ", len(s)) // 11 因為一個中文佔 3  bytes

	// str to int
	a, _ := strconv.Atoi("123")
	fmt.Println("str to int = ", a)

	// int to str
	b := strconv.Itoa(123)
	fmt.Println("int to str = ", b)

	// str to byte
	bytes := []byte("hello")
	fmt.Println("str to byte = ", bytes)
	fmt.Printf("str to byte = %c", bytes)

	// byte to string
	//s1 := string([]byte{})

}
