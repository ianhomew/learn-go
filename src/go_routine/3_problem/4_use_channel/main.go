package main

import "fmt"

type Cat struct {
	Name string
	Age  int
}

func main() {

	var allChan chan interface{}
	allChan = make(chan interface{}, 5)

	// write
	allChan <- Cat{
		Name: "Tom",
		Age:  10,
	}
	allChan <- 10
	allChan <- 20.33
	allChan <- "CAT STRING"
	allChan <- true

	// read
	cat := <-allChan
	// 注意: 出來是原本的類型 但是還是要轉型...編譯器無法識別 因為是 interface {}
	fmt.Printf("type = %T, cat = %v\n", cat, cat)
	// fuck 不能這樣用 要斷言才可以
	//fmt.Println("cat name = ", cat.Name)
	fmt.Println("cat name = ", cat.(Cat).Name)
	fmt.Println(<-allChan)
	fmt.Println(<-allChan)
	fmt.Println(<-allChan)

}
