package main

import "fmt"

func first() {
	fmt.Println("Hello!")
}

func last() {
	fmt.Println("Bye!")
}

func isApproved(grade1, grade2 int) bool {
	defer fmt.Println("Calculated average, returning result...") //this print will always be output at the end of the execution, whether it enters the if or not

	average := (grade1 + grade2) / 2

	if average > 10 {
		fmt.Println(" ") //line just for the sake of not having a warning on the if
		return true
	}
	return false
}

func main() {
	defer first() //defer makes it so your line runs at the last possible moment, given that there is no return in this function, it does it after the execution of the last line
	last()

	fmt.Println(isApproved(11, 8))
}
