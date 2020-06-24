package main

import "fmt"

type Point struct {
	x, y int
}

func main() {

	var emptyInterface interface{}
	var point Point = Point{1, 2}

	emptyInterface = point                                                               // ok
	fmt.Printf("emptyInterface type = %T, value = %v\n", emptyInterface, emptyInterface) // emptyInterface type = main.Point, value = {1 2}

	var b Point

	// b = emptyInterface // error
	b = emptyInterface.(Point) // 類型斷言
	fmt.Printf("b type = %T, value = %v\n", b, b)

	var f float32 = 333.111
	emptyInterface = f
	var ff float32
	// 斷言一定要相同的資料類型
	ff = emptyInterface.(float32) // 類型斷言 感覺斷言就是強制轉型？
	//ff = emptyInterface.(float64) // 這樣會失敗

	fmt.Printf("ff type = %T, value = %v\n", ff, ff)

	// 如何檢測
	var i int = 99
	emptyInterface = i
	fmt.Println("emptyInterface = ", emptyInterface)

	// 標準
	j, ok := emptyInterface.(int)
	if ok {
		fmt.Println("j=", j)
	} else {
		fmt.Println("ok = ", ok)
	}

	// 簡潔 if (宣告); 判斷式
	if j, ok := emptyInterface.(int); ok {
		fmt.Println("j=", j)
	} else {
		fmt.Println("ok = ", ok)
	}

}
