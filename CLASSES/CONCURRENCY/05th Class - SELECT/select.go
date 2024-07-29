package main

import (
	"fmt"
	"time"
)

func main() {
	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Channel 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			channel2 <- "Channel 2"
		}
	}()

	for {
		/* INCORRECT EXAMPLE, UNCOMMENT AND COMMENT BELOW TO SEE IT WORK
		messageChannel1 := <-channel1 //this line is ready to receive a value every 0.5 seconds
		fmt.Println(messageChannel1)

		messageChannel2 := <-channel2 //but it has to wait for this line because it only receives values every 2 seconds, this generates an unnecessary delay
		fmt.Println(messageChannel2) */

		select { //using select, we no longer have to have any lines waiting for each other, because the loop will select whichever case it needs depending on what's ready to be received
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}

}
