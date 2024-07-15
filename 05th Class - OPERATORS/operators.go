package main

import "fmt"

func main() {
	//ARITHMETIC
	sum := 1 + 1
	sub := 2 - 1
	mult := 2 * 2
	div := 4 / 2
	remainder := 11 % 2

	fmt.Println(sum, sub, mult, div, remainder)

	//you may not apply operands to two different types, so for example int and float, or even int32 and int64
	//this is because go is very strongly typed

	//RELATIONAL

	equal := (1 == 1)
	greater := (2 > 1)
	diff := (1 != 2)

	fmt.Println(equal, greater, diff)

	//LOGIC

	and := (true && false)
	or := (true || false)
	neg := !true

	fmt.Println(and, or, neg)

	//UNARY

	number := 1

	number++ // += also works
	println(number)
	number -= 1
	println(number)
	number *= 4
	println(number)

	// /= and %= also work

	//TERNARY OPERATORS DO NOT EXIST IN GO

}
