package typewriter

import "fmt"

func bye() { //functions in lowercase are of private access and can only be accessed in their own file
	fmt.Println("Bye world!")
}

func Bye() {
	bye()
}
