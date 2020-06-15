package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var s = "hello世界"

	// 1. string len
	// len can use to count array size
	fmt.Println("s len = ", len(s)) // 11 因為一個中文佔 3  bytes

	// 2. loop string
	// has problem: loop string and deal with Chinese word
	for i := 0; i < len(s); i++ {
		// has problems, because a Chinese word has three bytes
		fmt.Printf("the word is %c\n", s[i])
	}
	// no problem
	var s2 = []rune(s)
	for i := 0; i < len(s2); i++ {
		fmt.Printf("the word is %c\n", s2[i])
	}

	// 3. str to int
	a, _ := strconv.Atoi("123")
	fmt.Println("str to int = ", a)

	// 4. int to str
	b := strconv.Itoa(123)
	fmt.Println("int to str = ", b)

	// 5. str to byte
	bytes := []byte("hello") // 給一個字串 強制轉成 byte slice
	fmt.Printf("str to byte = %c\n", bytes)
	fmt.Printf("str to byte = %v\n", bytes)

	// 6. byte to string
	s1 := string(bytes)
	fmt.Println("byte to str = ", s1)
	var s3 = string([]byte{97, 98, 99})
	fmt.Println("byte to str = ", s3)
	fmt.Println("byte to str = ", s3)

	// 7. 10's to 2's,   10's to 8's,    10's to 16's
	var _10s int64 = 256
	fmt.Println("10 to 2", strconv.FormatInt(_10s, 2))
	fmt.Println("10 to 8", strconv.FormatInt(_10s, 8))
	fmt.Println("10 to 16", strconv.FormatInt(_10s, 16))

	// 8. find substring in string
	var superMan = "I am superman"
	if strings.Contains(superMan, "super") {
		fmt.Println(superMan, "has super")
	}

	// 9. 統計一個字串中含有幾個指定的子字串
	var cheese = "cheese"
	fmt.Println(cheese, "has", strings.Count(cheese, "e"), "e")   // 3
	fmt.Println(cheese, "has", strings.Count(cheese, "ee"), "ee") // 1

	// 10. 不區分大小寫的字串比較, 區分大小寫 直接使用 ==
	if strings.EqualFold("apple", "APPLE") {
		fmt.Println("apple is equal to APPLE")
	}
	fmt.Println("apple is not equal to APPLE,", "apple" == "APPLE")

	// 10. get the index of substring in the string, if none then -1
	var s4 = "ABCDEFG"
	fmt.Println(s4, ", the index of B is", strings.Index(s4, "B")) // 1
	fmt.Println(s4, ", the index of B is", strings.Index(s4, "Z")) // -1
	//strings.LastIndex()

}
