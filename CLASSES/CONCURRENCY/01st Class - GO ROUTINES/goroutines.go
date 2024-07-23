package main

import (
	"fmt"
	"time"
)

func main() {
	go write("Hello World!") //goroutine (go <CODE_LINE>) makes the code run past the line of code without waiting for its end, meaning write("Programming Go!") will start right after the program goes into this line
	write("Programming Go!")
	//if i was to put a go routine in the line above, the program would end without outputting anything because it would reach the ending of main before even running anything
}

func write(str string) {
	for {
		fmt.Println(str)
		time.Sleep(time.Second)
	}
}
