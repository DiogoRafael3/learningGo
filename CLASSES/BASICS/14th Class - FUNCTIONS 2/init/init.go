package main

import "fmt"

var n int

func init() { //init is performed before the execution of the main function, it is exactly like a @postconstruct if you've used spring before
	fmt.Println("Printing stuff...")
	n = 1
}

func main() {
	fmt.Println("Main stuff...", n)
}
