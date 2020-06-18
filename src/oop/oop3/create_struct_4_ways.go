package main

import "fmt"

// 結構體是值類型
type Person struct {
	Name   string
	Age    int
	Scores [5]float64
	ptr    *int
	slice  []int
	map1   map[string]string
}

func main() {

	// 1.
	var p1 Person
	fmt.Println(p1)

	// 2.
	var p2 Person = Person{}
	fmt.Println(p2)

	// 3.
	// 是一個指針
	var p3 *Person = new(Person)
	// 因為 p3 是指針 因此標準的給值方式如以下
	(*p3).Name = "Person3"
	(*p3).Age = 30
	fmt.Println(*p3)

	// 以上寫法好痛苦 go 很友善．．．可以這樣寫
	p3.Name = "P3"
	p3.Age = 40
	// 可以這樣用 是因為 go 在底層有做處理 會給 p3 加上「取值運算」
	// 所以 p3 底層幫你變成 *p3

	fmt.Println(p3) // &{P3 40 [0 0 0 0 0] <nil> [] map[]}

	// 4.
	var p4 *Person = &Person{}
	(*p4).Age = 20
	p4.Name = "Allen"

	fmt.Println(p4)

}
