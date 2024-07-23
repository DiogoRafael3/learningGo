package main

import "fmt"

func calc(n1, n2 int) (sum int, sub int) { //i can name the return of the variables using (<VARNAME_1> <VARTYPE_1>, ... , <VARNAME_N> <VARTYPE_N>)
	sum = n1 + n2
	sub = n1 - n2
	return // this automatically returns both variables as i named them at the declaration of the function
}

func main() {

	sum, sub := calc(10, 20)

	fmt.Println(sum, sub)
}
