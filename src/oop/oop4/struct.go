package main

import "fmt"

// 結構體是值類型
// 結構體的記憶體位址分配 是連續的
type Person struct {
	Name string
	Age  int
}

func main() {

	var p1 = Person{}
	p1.Name = "Person1"

	var p2 = &p1
	p2.Name = "Changed Name"

	fmt.Println(p1.Name)
	fmt.Println(p2.Name)

	fmt.Printf("%p\n", &p1) // 0xc00000c060

	fmt.Printf("%p\n", p2)  // 值 0xc00000c060
	fmt.Printf("%p\n", &p2) // 記憶體位址 0xc00000e028

}
