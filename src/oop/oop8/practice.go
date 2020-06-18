package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) speak() {
	fmt.Println(p.Name, "是個好人")
}

// 可以進行運算
func (p Person) cal() {
	res := 0
	for i := 1; i < 10000; i++ {
		res += i
	}
	fmt.Println(p, " cal() result = ", res)
}

func (p Person) cal2(n int) {
	res := 0
	for i := 1; i < n; i++ {
		res += i
	}
	fmt.Println(p, " cal2() result = ", res)
}

// 這邊參數是 值拷貝 也可以用指針
// 值拷貝的問題在於 我們把相同的東西複製一份 佔用了記憶體空間
func (p Person) getSum(a, b int) int {
	return a + b
}

// 給 Person 添加 speak 方法 輸出 xxx 是個好人
func main() {

	var p Person = Person{"Tom"}

	p.speak()
	p.cal()
	p.cal2(10)
	res := p.getSum(1, 2)
	fmt.Println("1 + 2 = ", res)
}
