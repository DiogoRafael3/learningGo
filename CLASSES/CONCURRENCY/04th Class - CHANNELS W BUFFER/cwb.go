package main

import "fmt"

func main() {
	channel := make(chan string, 2) //this specifies a capacity for my channel, allowing it to wait to receive two values (buffer), the default buffer for a channel is 0
	//meaning there should already be something ready to receive a value when a send is performed

	//send is a blocking operation, which is why we used a different function last class, to elaborate this will expect a channel to consume the string and since there is nothing
	//currently ready to receive it, it will generate a deadlock as the program will wait forever, this is the reason for giving the channel a capacity
	channel <- "Hello World" //this is a static send operation
	channel <- "Programming" //now we give a second value to the channel and it waits to receive it because we gave it a buffer of 2
	// if a third value is sent, a deadlock is created

	message := <-channel
	message2 := <-channel //receive the second value given to the channel
	//if a third value is waiting to receive something, a deadlock is caused

	fmt.Println(message)
	fmt.Println(message2) //this print could even not be used, however the second string would go to waste and remain unused
}
