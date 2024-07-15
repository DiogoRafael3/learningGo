package main

import "fmt"

func main() {

	function_result := func(text string) string {
		return fmt.Sprintf("Anonymous %s", text) //sprintf creates a string based on what you give it
	}("Parameter") //anonymous funcion, as soon as it is declared it is executed

	fmt.Println(function_result)
}
