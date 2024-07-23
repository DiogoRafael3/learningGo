package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wait_group sync.WaitGroup //sync.WaitGroup is from the base go package, this is a way to wait for multiple goroutines to finish their execution

	wait_group.Add(2) //add the number of goroutines to the wait group

	go func() { //anonymous function created to use wait_group.Done()
		write("Hello World!")
		wait_group.Done() //this performs a -1 in the wait_group
	}()

	go func() {
		write("Programming Go!")
		wait_group.Done() //when both wait groups are done, then the program may proceed
	}()

	wait_group.Wait() //this marks where the program waits for the wait groups

	fmt.Println("All done!") //this will only be done after both goroutines are done
}

func write(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
		time.Sleep(time.Second)
	}
}
