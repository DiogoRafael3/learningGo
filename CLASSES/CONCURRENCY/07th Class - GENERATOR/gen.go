package main

import (
	"fmt"
	"time"
)

func main() {
	channel := write("sup")

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func write(text string) <-chan string { //this is the generator pattern, it is used to abstract the creation of complex goroutines in the main function
	channel := make(chan string)

	go func() { //goroutine that we don't want to write in main, could be anything, such as the fibonacci sequence for example
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return channel
}
