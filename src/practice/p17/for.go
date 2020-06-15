package main

import "fmt"

func main() {

	// 1.
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// 2.
	var j = 0
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 3.
	var k = 0
	for {
		fmt.Println(k)
		if k < 10 {
			break
		}
		k++
	}

	// 4. for
	var s = "ABCDEFG中文"
	//for i:= 0; i<len(s); i++ {
	//	fmt.Printf("%c", s[i]) // 中文佔 3 bytes 英文字母佔 1 byte 會出現亂碼
	//}

	// 上面的解法
	var s2 = []rune(s)
	for i := 0; i < len(s2); i++ {
		fmt.Printf("=== %c", s2[i]) // 這樣不會出現亂碼
	}

	// 5. for range
	for index, value := range s {
		// 不會有問題 注意 index 的值 會發現 index 從7 跳到 10 因為中文佔 3 bytes
		fmt.Printf("index = %d, value = %c\n", index, value)
	}

}
