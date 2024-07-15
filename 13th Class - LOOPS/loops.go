package main

import (
	"fmt"
	"time"
)

func main() {
	//GO ONLY HAS "FOR" FOR LOOPS
	i := 0

	for i < 3 { // for <BOOLEAN>
		time.Sleep(time.Second) //using time library just to watch the loop go in the console, comment this to see the result instantly
		i++
		fmt.Println(i)
	}

	//you may also use inits inside for declarations, however they are restricted to the scope as well
	for x := 0; x > -3; x -= 2 { // you may iterate however many at a time
		fmt.Println(x)
	}

	names := [3]string{"John", "Azores", "Portugal"}

	for _, name := range names { //this is a for each in go, the first var refers to the position of the iterable and the second refers to the value, by using _ we forgo the assignment of index to a variable
		fmt.Println(name)
	}

	for index, letter := range "WORD" { //as previously said, there is no chars in go so this outputs the ascii code in each position rather than the letter
		fmt.Println(index, string(letter)) //by using string(<ASCII_CODE>) we convert the ascii code into string, this is a cast
	}

	//ITERATING OVER MAPS

	user := map[string]string{
		"name":    "Daniel",
		"surname": "Cuts",
	}

	for key, value := range user { //this iterates over maps in order
		fmt.Println(key, value)
	}

	//to create an infinite loop, just do a for{} without a condition
}
