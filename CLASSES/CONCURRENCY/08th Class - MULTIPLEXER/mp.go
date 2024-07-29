package main

import (
	"fmt"
	"time"
)

func main() {
	//using the multiplexer pattern, we can use two different entry channels and send their results into the same output channel in main
	channel := multiplex(write("dog", 1000), write("cat", 2000))

	//DO NOTE THAT THE ORDER YOU WRITE THE CHANNELS IN DOES NOT AFFECT WHICH GOES FIRST
	//this means that in this program, either cat or dog could go first in the command line

	for i := 0; i < 10; i++ {
		fmt.Println(<-channel)
	}
}

func multiplex(entryChannel1, entryChannel2 <-chan string) <-chan string { //multiplexer pattern
	outputChannel := make(chan string) //channel we're using to merge both entry channels into one

	go func() { //picks between the two channels we sent to the function as arguments
		for {
			select {
			case message := <-entryChannel1:
				outputChannel <- message
			case message := <-entryChannel2:
				outputChannel <- message
			}
		}
	}()

	return outputChannel
}

func write(text string, ms int) <-chan string { //generator
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Value received: %s", text)
			time.Sleep(time.Millisecond * time.Duration(ms))
		}
	}()

	return channel
}
