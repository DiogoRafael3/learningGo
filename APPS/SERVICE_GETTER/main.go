package main

import (
	"log"
	"os"
	"serv_get/app"
)

func main() {

	application := app.Generate()

	if err := application.Run(os.Args); err != nil { //this function allows us to run the application while getting the args from the command line, it outputs an error so we can treat it immediatly
		log.Fatal("Failed to start app: ", err)
	}

}
