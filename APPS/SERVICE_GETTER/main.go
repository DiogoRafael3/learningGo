package main

import (
	"log"
	"os"
	"serv_get/app"
)

func main() {

	//to run this application either run
	//go run main.go ip --<COMMAND> <WEB_ADRESS>
	//or
	// ./serv_get ip --<COMMAND> <WEB_ADRESS>
	//ip being the name of the command
	application := app.Generate()

	if err := application.Run(os.Args); err != nil { //this function allows us to run the application while getting the args from the command line, it outputs an error so we can treat it immediatly
		log.Fatal("Failed to start app: ", err)
	}

}
