package main

import "fmt"

func main() {
	tasks := make(chan int, 45) //creating a channel to receive 45 tasks
	results := make(chan int, 45)

	//you can use multiple of these worker pools to speed up the process, instead of running one fibonacci at a time, you can run multiple, this is because each task
	//of the channel is done in order, making it so the channels know which fibonacci number to calculate and which result to send to the results channel
	go worker(tasks, results)
	go worker(tasks, results)
	go worker(tasks, results)
	//go worker(tasks, results)
	//the more worker pools you use the faster the program, however this is limited to your computer's processing power, higher number might not mean significant increases if
	//your computer can't handle running a lot of processes at the same time

	for i := 0; i < 45; i++ {
		tasks <- i
	}
	close(tasks)

	for i := 0; i < 45; i++ {
		result := <-results
		fmt.Println(result)
	}
}

func worker(tasks chan int, results chan int) {
	for number := range tasks { //this waits for tasks to arrive, therefore you can call this function using a goroutine and it'll wait for the channel to be populated before trying to run values
		results <- int(fibonacci(uint(number))) //this will start sending values to results as tasks arrive
	}
}

func fibonacci(pos uint) uint { //we use the fibonacci function with high numbers to show how effective worker pools are
	if pos <= 1 {
		return pos
	}

	return fibonacci(pos-2) + fibonacci(pos-1) //recursively calling itself
}
