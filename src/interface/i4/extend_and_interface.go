package main

import "fmt"

// 定義接口
type Birdable interface {
	flying()
}
type Fishable interface {
	swimming()
}

type Monkey struct {
	Name string
}

func (monkey *Monkey) climbing() {
	fmt.Println(monkey.Name, "爬樹")
}

type LittleMonkey struct {
	Monkey
}

// 悟空要會飛！
// 這樣只有悟空可以飛
func (lm *LittleMonkey) flying() {
	fmt.Println(lm.Name, " flying")
}

// 悟空要會飛！
// 這樣只有悟空可以飛
func (lm *LittleMonkey) swimming() {
	fmt.Println(lm.Name, "  swimming")
}

func main() {

	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}

	// 爬樹 是猴子都會的 所以是 func (monkey *Monkey) climbing()
	monkey.climbing()
	// 飛是特殊的 只有悟空這隻猴子會 所以定義一個介面 讓小猴子這個結構體去實作「飛」
	// 這樣不會破壞原本猴子的屬性與能力(因為普通猴子不會飛跟游泳) 如果用繼承就會
	monkey.flying()
	monkey.swimming()

	// 當Ａ結構體需要「擴充功能」 同時不希望去破壞繼承的關係 則可以去實現某個接口即可
	// 因此可以認為 實現接口是對繼承機制的補充

}
