package main

import "fmt"

func main() {

	fmt.Println("請輸入 A, B, C, D")

	var abcd byte
	_, _ = fmt.Scanf("%c", &abcd)
	fmt.Printf("your input is %c\n", abcd)

	// switch 的 abcd, 'A', 'B', 'C', 'D' 資料類型必須相同
	switch abcd {
	case 'A':
		fmt.Println("等級A")
	case 'B':
		fmt.Println("等級B")
	case 'C':
		fmt.Println("等級C")
	case 'D':
		fmt.Println("等級D")
	default:
		fmt.Printf("your input is %c\n", abcd)
	}

}
