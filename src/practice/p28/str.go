package main

import (
	"fmt"
	"strings"
)

func main() {

	// 12. find substring at the last position
	var s1 = "go golang"
	fmt.Println("last index = ", strings.LastIndex(s1, "go"))

	// 13. replace n: 替換次數 -1表示全部
	var s2 = "go golang"
	fmt.Println(strings.Replace(s2, "go", "語言", 1))  // 語言 golang
	fmt.Println(strings.Replace(s2, "go", "語言", 2))  // 語言 語言lang
	fmt.Println(strings.Replace(s2, "go", "語言", -1)) // 語言 語言lang

	// 14. 將字串拆分成陣列, in php is explode
	var s3 = "hello,b,c,d,e,f,中文"
	var chars = strings.Split(s3, ",")
	for i := 0; i < len(chars); i++ {
		fmt.Println("chars = ", chars[i])
	}

	// 14. to upper case
	// 15. to lower case
	var s4 = "GO LANG"
	var s5 = "go lang"
	fmt.Println(s4, "to lower = ", strings.ToLower(s4))
	fmt.Println(s5, "to upper = ", strings.ToUpper(s5))

	// 16. 去掉左右空格
	var s6 = "        AAA BBB         "
	fmt.Printf("%q\n", strings.TrimSpace(s6))

	// 17. 去掉左右兩邊指定的字符
	var s7 = " !  AA!!A     B!BB   CC!!C   "
	fmt.Printf("%q\n", strings.Trim(s7, " !A"))
	fmt.Printf("%q\n", strings.Trim(s7, "A! "))

	// 18. 只去掉左邊
	var s8 = " !  AA!!A     B!BB   CC!!C!!!   "
	fmt.Printf("%q\n", strings.TrimLeft(s8, "! A"))

	// 19. 只去掉右邊
	fmt.Printf("%q\n", strings.TrimRight(s8, "! A"))

	// 20. has prefix
	var s9 = "http://localhost:8080"
	fmt.Println("has http:// = ", strings.HasPrefix(s9, "http://"))

	// 21. has suffix
	fmt.Println("has 8080 = ", strings.HasSuffix(s9, "8080"))

}
