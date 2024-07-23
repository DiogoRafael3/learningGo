package main

import "fmt"

func dayOfTheWeek(number int) string {
	//SWITCHES IN GO DO NOT NEED BREAKS
	switch number { //switch <ARGNAME> {case <BOOLEAN>:}
	case 1:
		return "Sunday"
	case 2:
		return "Monday"
	case 3:
		return "Tuesday"
	case 4:
		return "Wednesday"
	case 5:
		return "Thursday"
	case 6:
		return "Friday"
	case 7:
		return "Saturday"
	default:
		return "Invalid number"
	}
	//If no case is valid and no default is set, then you need to use a return after the switch
}

func madness(number int) string {
	var gauge string

	switch {
	case 0 < number && number <= 30:
		gauge = "Not mad at all"
	case 30 < number && number <= 70:
		gauge = "A bit mad"
	case number > 70:
		gauge = "TOTALLY MAD!"
		fallthrough //fallthrough makes you fall into the next case without even checking the condition, does seem kind of niche though
	case number < 0:
		gauge = "ABSOLUTELY CRAZY!!!!"
	}

	return gauge
}

func main() {
	fmt.Println(dayOfTheWeek(0))
	fmt.Println(dayOfTheWeek(2))

	fmt.Println(madness(71)) //outputs the last case because there was a fallthrough on the penultimate case
}
