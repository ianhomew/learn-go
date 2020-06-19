package main

type A struct {
	Name string
	int  // 重點是這個
}

func main() {

	a := A{}
	a.int = 99 // 可以直接這樣用

}
