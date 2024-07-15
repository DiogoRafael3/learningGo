package main

import "fmt"

var n int

func init() { //init is like a postconstruct in spring boot, it is performed before the execution of the main function
	fmt.Println("Printing stuff...")
	n = 1
}

func main() {
	fmt.Println("Main stuff...", n)
}
