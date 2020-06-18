package main

import "fmt"

type Pupil struct {
	Name  string
	Age   int
	Score int
}

func (p *Pupil) ShowInfo() {
	fmt.Printf("學生 %v, 年齡 %v, 成績為 %v", p.Name, p.Age, p.Score)
}
func (p *Pupil) SetScore(score int) {
	p.Score = score
}
func (p *Pupil) Testing() {
	fmt.Println("小學生正在考試中...")
}

func main() {

	var p = &Pupil{
		"tom",
		10,
		0,
	}

	p.Testing()
	p.SetScore(90)
	p.ShowInfo()

}
