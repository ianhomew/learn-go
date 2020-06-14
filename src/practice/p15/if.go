package main

import "fmt"

func main() {

	fmt.Println("請輸入一到一百的數字")

	var num int
	_, _ = fmt.Scanf("%d", &num)
	fmt.Println("your input is ", num)

	if num <= 10 {
		fmt.Println("小於等於１０")
	} else if num >= 11 && num <= 60 {
		fmt.Println("11-60 之間")
	} else {
		fmt.Println(" >= 61")
	}

}
