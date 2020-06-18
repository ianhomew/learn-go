package main

import "fmt"

// 方法 是作用在指定的資料型別上(和指定的資料類型綁定)
// 因此自定義類型 都可以有方法 不只是結構體 int float32 都可以

// 函數與方法的差異
// 調用方式不同
// 函數: 函數名(參數...)
// 方法: 變數.方法名(參數...)
//
// 普通函數如果參數是傳值 不能將指標類型的資料進行傳遞 反之亦然
// 對於方法(ex. struct 的方法) 接收者為值類型時 可以直接用指標類型的變數調用方法 反過來也可以
//
//
//

type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}
func (i *integer) change() {
	*i = *i + 1

}

type Student struct {
	Name string
	Age  int
}

// 如果一個類型 實作了 String() 方法 則 fmt.Println 會調用該 String()方法
func (s *Student) String() string {
	return "自定義 String() 寫log 好用"
}

func main() {

	var i integer = 10
	i.print()

	i.change()
	fmt.Println("new i = ", i)

	var student = Student{"Tom", 13}
	fmt.Println(student)
	fmt.Println(&student)

}
