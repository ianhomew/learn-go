package main

import "fmt"

type Goods struct {
	Price float64
	Name  string
}
type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

func main() {

	tv := TV{
		Goods: Goods{
			Price: 99.99,
			Name:  "電視001",
		},
		Brand: Brand{
			Name:    "Sharp",
			Address: "Taipei",
		},
	}
	fmt.Println(tv)

	tv2 := TV2{
		Goods: &Goods{
			Price: 100,
			Name:  "指標電視",
		},
		Brand: &Brand{
			Name:    "指標品牌名稱",
			Address: "指標地址",
		},
	}

	fmt.Println(tv2.Brand.Name)
	fmt.Println(tv2.Brand)

}
