package main

import "fmt"

func main() {

	var var1 int = 10
	var var2 int = var1 //when you assign a var to a  value, it assigns a copy of the var, it does not create a relation between the two vars

	fmt.Println(var1, var2) //outputs 10 and 10

	var1++

	fmt.Println(var1, var2) //outputs 11 and 10

	//a pointer is a memory reference

	var var3 int
	var pointer *int

	fmt.Println(var3, pointer) //because it is a memory reference and not a number in itself, pointer will be <nil> and not 0 in this print

	var3 = 100
	pointer = &var3 //this now assings the reference of var3 to the pointer

	fmt.Println(var3, pointer)
	fmt.Println(*pointer) //this consumes the reference and goes to the memory adress to get the value

	var3++

	fmt.Println(var3, *pointer) //now both values should be 101 as we have added 1 to the value contained in the memory adress

}
