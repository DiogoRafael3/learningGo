package main

import "fmt"

type user struct { //structs are like objects in java, they are a type that can contain many different fields
	name   string
	age    uint8
	adress adress //you can also use structs inside of structs, and you may name your field the same as the struct name
}

type adress struct { //unlike c, code is not compiled strictly in order so you may put your secondary structs above or below your main struct
	street string
	number uint8
}

func main() {

	//to use field assignment in go, simply do <VARNAME>.<FIELDNAME> = <VALUE>
	var u user
	u.name = "Diogo"
	u.age = 25
	fmt.Println(u)

	//if you wish to declare the whole struct at once then use <TYPENAME>{<FIELDS>}
	adress := adress{"Rua QJSC", 41}
	user2 := user{"David", 38, adress}
	fmt.Println(user2)

	//to declare only one field in the struct, use <TYPENAME>{<FIELD>:<VALUE>}
	user3 := user{name: "Rita"}
	user3.age = 40
	fmt.Println(user3)

}
