package main

import "fmt"

func closure() func() {
	text := "Closure function text"

	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	//this should be obvious to anyone that has programmed for a while but the text declared inside closure will always be the text declared within the scope of the function
	text := "Main function text"
	fmt.Println(text)

	newFunction := closure()
	newFunction() //this will always print "Closure function text" no matter what, despite having a text variable declared before it

}
