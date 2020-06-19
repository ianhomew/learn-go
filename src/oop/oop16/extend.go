package main

import "fmt"

type Person struct {
	Name string
}

// 在 Student 結構體內 使用了匿名結構體 Person
type Student struct {
	Person
	Name string
}

func (student *Student) SayName() {
	fmt.Println("student say", student.Name)
}
func (person *Person) PrintPersonName() {
	fmt.Println("student say", person.Name)
}

func main() {

	var student Student

	// 因為 student 本身的有 Name 屬性
	// 所以不會 把值給到 Person 內的 Name
	student.Name = "Bob"

	fmt.Println("student.Name = \"Bob\"")
	fmt.Printf("student name is %q\n", student.Name)
	fmt.Printf("person name is %q\n", student.Person.Name)

	student.SayName()

	// 因為沒有把值給到 Person 結構體的 Name 所以印出來的是預設值 「空」
	student.PrintPersonName()

}
