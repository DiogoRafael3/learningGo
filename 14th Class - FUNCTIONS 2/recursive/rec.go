package main

import "fmt"

func fibonacci(pos uint) uint {
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1) //recursively calling itself
}

func main() {
	fmt.Println(fibonacci(10)) //recursive functions are extremely slow when dealing with big numbers because going into a function everytime creates a big stack which may overflow

	//fibonacci sequence print
	for i := uint(0); i <= 10; i++ {
		fmt.Println(fibonacci(i))
	}

}
