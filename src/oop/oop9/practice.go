package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// 為了提高效率 綁定給結構體的方法 會用 pointer -> c *Circle 參數就不一定
// 通常都傳指標過來 效率比較好
func (c *Circle) areaPointer() float64 {
	//return 3.14 * c.Radius * c.Radius
	return 3.14 * (*c).Radius * (*c).Radius
}

func main() {

	c := Circle{4.0}
	res := c.area()
	fmt.Println("res = ", res)
	res1 := c.areaPointer()
	fmt.Println("res1 = ", res1)

	// 標準使用方式 因為 (c *Circle) c 是一個指標
	//res1 := (&c).areaPointer()

}
