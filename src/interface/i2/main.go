package main

import (
	"fmt"
)

// 1. 介面/接口(interface) 不能建立實體 但是可以指向一個實現了該接口的自定義類型的變數
// 2. 接口中所有的方法都沒有真正實作
// 3. 在 go 中 一個自定義類型需要將某個接口的所有方法都實現 我們說這個自定義類型實現了該接口
// 4. 一個自定義類型只有實現了某個接口 才能將該自定義類型的實體賦給接口類型 (就是強制轉換)
// 5. 只要是自定義的資料類型 都可以實現接口 不僅僅是結構體類型
// 6. 一個自定義類型可以實現多個接口
// 7. go 中接口不能有任何變數
// 8. 一個接口(A)可以繼承多個別的接口(B and C) 如果要實作 A 接口 則必須實作B, C 接口
// 9. interface 默認是指針類型 如果沒有對 interface 初始化就使用 會輸出 nil
// 10. 空接口 「interface {}」沒有任何方法 所以「可以把任何變數賦給空接口」

// 10. EX: var i interface{} = Student

type AInterface interface {
	Say()
}

type integer int

func (i *integer) Say() {
	fmt.Println("integer say i = ", *i)
}

type T interface {
}

func main() {

	var i integer = 100
	i.Say()

	var j integer = 222
	//fmt.Println("j size = ",unsafe.Sizeof(j))
	var cast AInterface = &j
	//fmt.Println("cast size = ",unsafe.Sizeof(cast))
	cast.Say()

	var t T = i
	fmt.Println(t)
	var t2 interface{} = i // 這使用的非常頻繁 目前不知道用處
	fmt.Println(t2)        // 這樣 t2 就不會有 Say()方法了啊...

	var num1 float64 = 33.44
	var t3 = num1
	fmt.Println(t3)

}
