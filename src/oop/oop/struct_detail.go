package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

type Stu Student

func main() {

	var student1 Student

	var student2 Stu

	//student2 = student1 // 不可以
	//student2 = Stu(student1) 可以

	fmt.Println(student1, student2)

}
