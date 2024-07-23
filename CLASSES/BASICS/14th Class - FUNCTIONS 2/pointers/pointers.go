package main

import "fmt"

func invert(number int) int {
	return number * -1
}

func invertWithPointer(number *int) { //receives the pointer to the number
	*number = *number * -1
}

func main() {
	number := 20
	inverted := invert(number) //this creates a copy of the number, that was then inverted and returned
	fmt.Println(inverted)

	reference := &number //even with a middleman variable, the variable to be inverted will be changed because this middleman contains the reference to the variable

	invertWithPointer(reference) //this inverts the number itself
	fmt.Println(number)

}
