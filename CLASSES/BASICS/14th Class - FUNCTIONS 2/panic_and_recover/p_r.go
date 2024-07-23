package main

import "fmt"

func recovering() {
	if r := recover(); r != nil { //using an if init, recover can recover the execution of a program after a panic
		fmt.Println("Recovering...")
	}
}

func isApproved(grade1, grade2 float64) bool {
	defer recovering() //we use recover here to prevent the panic from going off if possible
	average := (grade1 + grade2) / 2

	if average > 6 {
		return true
	} else if average < 6 {
		return false
	}

	panic("AVERAGE IS 6") //panic kills the execution of the program if it is executed, outputting the message in the args
}

func main() {
	fmt.Println(isApproved(6, 7))
	fmt.Println(isApproved(6, 5))
	fmt.Println(isApproved(6, 6)) //returns false because it is the default value of a boolean

}
