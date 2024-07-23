package main

import (
	"fmt"
	"time"
)

func main() { //channels are ways to send and receive information from one goroutine to another, they are strongly typed and thus require you to use a single type of value
	channel := make(chan string) //this creates a string channel, meaning it can only receive and output strings

	go write("Hello World", channel)

	fmt.Println("After write is declared")

	for { //this will produce a deadlock if the channel is not properly handled, this is not identified in compilation
		message, open := <-channel //this symbolises that the channel is waiting to receive a value, however as soon as this receives a single value the program will keep going
		if !open {                 //open representes when channel is open, when channel is closed we break the for loop
			break
		}
		fmt.Println(message) //this means this will output a single Hello World, because it is the first thing channel receives
	}

	/* ALTERNATIVELY we can use this
	for message := range channel {
		fmt.Println(message)
	}
	to the same effect of the above code but neater and with less chance to produce errors
	*/

	//if a value is never sent (for example when all go routines are finished) then the channel will wait forever, this creates a deadlock
}

func write(str string, channel chan string) {
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		channel <- str //this passes str into the channel
		time.Sleep(time.Second)
	}

	close(channel) //this closes our channel
}
