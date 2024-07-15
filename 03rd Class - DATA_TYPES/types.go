package main

import (
	"errors"
	"fmt"
)

func main() {
	//NUMBERS

	//int types are split between int8, int16, int32 and int64
	//var num int8 = 999999 fails because the number overflows under that var type
	//var num int -> if your computer is 64 bit architecture, it will create an int64, if it is 32 bit then an int32
	//var num uint = -2 -> uint is unsigned int so negative numbers will not work

	num := 1000000 //infers the base int type, which means it uses your computer's architecture
	fmt.Println(num)

	var num2 rune = 12345 //rune = int32
	fmt.Println(num2)

	var num3 byte = 123 //byte = uint8
	fmt.Println(num3)

	//for floating numbers, float32 and float64 are the only two types
	//var <VARNAME> float is invalid, you may not use float by itself

	floating := 5437852.234 //same case as infering int
	fmt.Println(floating)

	//STRINGS

	var str string = "text"
	fmt.Println(str)

	str2 := "text"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char) //there is no char type in go, so this will be inferred as the ASCII code for the characterx

	//UNASSIGNED
	//empty declarations of strings and ints will result in the default values, being " " and 0 respectively and not null

	var empty_text string
	fmt.Println(empty_text)

	var empty_number int
	fmt.Println(empty_number)

	//BOOLEAN

	var boolean bool = true
	fmt.Println(boolean)

	var boolean2 bool //defaults to false
	fmt.Println(boolean2)

	//ERROR
	//error is a type utilized for error handling in go

	var errorvar error //defaults to <nil>
	fmt.Println(errorvar)

	var errorvar2 error = errors.New("internal error") //correct way to declare an error
	fmt.Println(errorvar2)
}
