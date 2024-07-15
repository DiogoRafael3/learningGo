package main

import "fmt"

func main() {
	num := 10

	if num > 15 { //if condition (if <BOOLEAN_COND>)
		fmt.Println("Greater")
	} else { //else
		fmt.Println("Lesser")
	}

	//go does not let you do single line ifs, you always need to use {} indentation

	if other_num := 1; other_num > 0 { //you may declare variables before conditions, separating both by a ; HOWEVER it is only available in the scope of the if
		fmt.Println("Number is greater than 0")
	} else if other_num < 0 { //else ifs also work in go
		fmt.Println("Number is lesser than 0")
	} else {
		fmt.Println("Number is 0")
	} //you may not access other_num after the if

	if num1, num2 := 1, 2; num1 < num2 { //you may initialize many different variables in the same if
		fmt.Println("2 is bigger than 1")
	}

}
