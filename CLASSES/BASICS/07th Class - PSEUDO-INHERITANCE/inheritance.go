package main

import "fmt"

type person struct {
	name string
	age  uint8
}

type student struct {
	person //declaring a struct type but not assigning it to a field makes it so all its fields are used in your current struct
	course string
	school string
}

func main() {

	var studnt student
	studnt.name = "Diogo" //student will now be able to access name and age from person without having to use student.person.name
	studnt.age = 25
	studnt.course = "Go"
	studnt.school = "Udemy"
	fmt.Println(studnt)

	studnt2 := student{person{"Miguel", 26}, "Go", "Udemy"} //the person struct must be declared if you use this type of assignment
	fmt.Println(studnt2)
}
