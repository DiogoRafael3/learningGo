package honorifics

import "strings"

// Verifies if your honorific is of a valid type following the format (<HONORIFIC> <FIRST NAME> <LAST NAME>)
func HonorificType(name string) bool {
	validTypes := []string{"mister", "miss", "mx"}

	honorific := strings.ToLower(strings.Split(name, " ")[0]) //lower case the first element of the string

	for _, typeIterate := range validTypes { //search for honorific in pre-determined array
		if typeIterate == honorific {
			return true
		}
	}

	return false
}
