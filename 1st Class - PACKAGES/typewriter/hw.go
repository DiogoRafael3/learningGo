package typewriter

import "fmt"

func Hello() {
	fmt.Println("Hello World!")
}

//this function is now accessible across the whole package
//code in this package must be under the directory that contains the go.mod file, be it in the same folder, or under another folder
