package model

type student struct {
	Name   string
	Score  float64
	score1 float64 // how to get this value in other .go file?
}

// 注意是 *student
// 放在堆裡面可以共享？
func NewStudent(name string, score float64) *student {
	// &student: 結構體本身是值類型 且該結構體為私有 外部無法取得
	// 所以返回一個指標 因為指標是另外一個記憶體空間？

	s := student{name, score, score}
	return &s
	// 上下一樣 寫成上面我暫時比較好理解................

	//return &student{
	//	name,
	//	score,
	//	score,
	//}
}
func (s *student) GetScore1() float64 {
	return s.score1
}
