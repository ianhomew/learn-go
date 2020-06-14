package main

import "fmt"

func main() {

	var name string
	var age int
	var sal float64
	var isPassTest bool

	// 方式一
	//fmt.Println("name:")
	//_, _ = fmt.Scanln(&name)
	//fmt.Println("age:")
	//age, _ = fmt.Scanln()
	//fmt.Println("sal:")
	//
	//// 怎麼用
	////sal,  _ = fmt.Scanf("%f", &sal)
	//fmt.Println("name = ", name, "age = ", age, "sal = ", sal)

	// 方式二
	_, _ = fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPassTest)
	fmt.Println("name = ", name, "age = ", age, "sal = ", sal, "isPassTest = ", isPassTest)

}
