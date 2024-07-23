package main

import (
	"fmt"
	"module/typewriter"
	//"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Initiating program...")
	typewriter.Hello()
	typewriter.Bye()

	/*error := checkmail.ValidateFormat("diogor@af8@gmail.com")
	fmt.Println(error)*/
}

//to run a single file:
//go run <FILENAME>
//to create an exe file with your current code:
//go mod init <APP_NAME>
//to refresh the application exe after having changed code:
//go build
//to import external packages, use the command:
//go get <PACKAGE_URL>
//when no longer using packages, use
//go mod tidy
