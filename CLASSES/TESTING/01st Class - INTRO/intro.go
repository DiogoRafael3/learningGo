package main

import (
	"fmt"
	hon "intro/honorifics" //this creates an alias for the honorifics called hon
)

// to test your whole file structure, on your root folder use -> go test ./...
// to see the result of every tests -> go test -v
// to check the coverage percentage of tests -> go test --cover
func main() {
	fmt.Println(hon.HonorificType("Mister Incredible"))
	fmt.Println(hon.HonorificType("Master Splinter"))

}
