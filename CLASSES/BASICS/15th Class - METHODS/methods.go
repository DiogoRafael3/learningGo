package main

import "fmt"

type user struct {
	name string
	age  uint8
}

//PLEASE NOTE THAT GO IS NOT OBJECT ORIENTED, DESPITE SUPPORTING METHODS
func (u *user) birthday() { //to declare a method : func (<STRUCTNAME> <STRUCT>) <FUNCNAME>()
	//the pointer was used above due to the value not changing if a pointer wasn't used, again, because a copy would be created and the value of that copy's age would be changed and not the user's itself
	fmt.Println("Inside birthday method...")
	u.age++
}

func (u user) isAdult() bool { //don't show this one to Drake
	return u.age >= 18 //to perform checks, save data into a database and such (basically anything that doesn't need changing the data) you don't have to use pointers
}

func main() {
	user1 := user{"John", 3}
	fmt.Println(user1)

	user1.birthday() //user ages one year
	fmt.Println(user1)

	fmt.Println(user1.isAdult())

}
