package main

import (
	"fmt"
)

func generic(interf interface{}) { //this is the usage of a generic in go (interface{}), because there is nothing inside the curly brackets, it will expect an interface of any kind
	fmt.Println(interf)
}

func main() {
	generic("i") //this function can be called with any type of interface, so we call it with string, int and boolean
	generic(1)
	generic(true)

	map1 := map[interface{}]interface{}{ //this should be used somewhat carefully as if it is misused it might end up too confusing, such as this map which uses 4 different types in different key value pairs
		1:            "String",
		float32(100): true,
		"String":     "String",
	}

	fmt.Println(map1) //note that this does work, however it is extremely confusing and not recommended
}
