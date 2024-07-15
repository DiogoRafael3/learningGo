package main

import "fmt"

func main() {
	var variable string = "Var1" //variables are declared under the notation var <VARNAME> <VAR_TYPE>
	variable2 := "Var2"          //however, go has type inference through :=

	fmt.Println(variable)
	fmt.Println(variable2)

	var ( //you may declare multiple variables at once using this notation
		variable3 string = "Var3"
		variable4 string = "Var4"
		//variable5 := "Var5" but type inference does not work under this scope
	)

	fmt.Println(variable3)
	fmt.Println(variable4)

	variable5, variable6 := "Var5", "Var6" //to declare multiple variables with type inference, this is the method that you have to use

	fmt.Println(variable5)
	fmt.Println(variable6)

	const constant string = "Const" //to declare constants, use the const keyword

	fmt.Println(constant)

	variable5, variable6 = variable6, variable5 //you may switch variable values like this

	fmt.Println("Variable 5 after switch:" + variable5)
	fmt.Println("Variable 6 after switch:" + variable6)
}
