package main

import "fmt"

func sumOfNNumbers(numbers ...int) (total int) { //by using ... we can obscure the number of parameters in a function, meaning you can send any number of ints to this function
	total = 0

	for _, number := range numbers {
		total += number
	}

	return
}

func write(text string, numbers ...int) { // you may also use it after any other variables, HOWEVER YOU MAY ONLY HAVE ONE VARIADIC VARIABLE PER FUNCTION, and it must be the last one in the function
	for _, number := range numbers {
		fmt.Println(text, number)
	}
}

func main() {
	fmt.Println(sumOfNNumbers(1, 2, 3, 4))
	write("i count to", 1, 2, 3)
}
