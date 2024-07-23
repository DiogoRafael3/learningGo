package main

import "fmt"

func main() {
	user := map[int]string{ //map is declared via map[<KEYTYPE>]<VALUETYPE>
		1: "Sophia",
		2: "Michael", //commas must be at the end of every element
	}

	fmt.Println(user)
	fmt.Println(user[1]) //accessing the key 1

	user[3] = "Beatrice" //to add an element to a map, simply declare it as <MAPNAME>[<KEY>] = <VALUE>

	fmt.Println(user)

	user2 := map[string]map[string]string{ //you can create maps that have maps as values
		"name": {
			"first":   "Peter",
			"surname": "Venture",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["name"])

	delete(user2, "name") //to delete items from a map, use the delete function
	fmt.Println(user2)

}
