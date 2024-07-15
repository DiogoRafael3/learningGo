package main

import "fmt"

func sum(n1 int8, n2 int8) int8 { //func <FUNCNAME>(<ARGS>) <RETURNTYPE>
	return n1 + n2
}

//if two params are of the same type and in succession, they may be declared only at the end
func calculus(n1, n2 int8) (int8, int8, string) { //functions may have multiple returns in go
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub, "This is crazy"
}

func main() {
	fmt.Println(sum(2, 3))

	sum := sum(30, 40) //functions are types, therefore may be inferred
	fmt.Println(sum)

	var f = func(txt string) string { //as they are types, they can be assigned
		fmt.Println(txt)
		return txt
	}

	f("Hello!") // you may call the function created within the scope of your current function

	result := f("This is America") //or you may assign it
	fmt.Println(result)

	fmt.Println(calculus(2, 1)) //prints will automatically assume the number of returns necessary

	resultSum, resultSub, text := calculus(5, 3) //however you must assign each return to a variable
	fmt.Println(resultSum, resultSub, text)

	result1, result2, _ := calculus(5, 5) //you may ignore a return by using underscore
	fmt.Println(result1, result2)
}
