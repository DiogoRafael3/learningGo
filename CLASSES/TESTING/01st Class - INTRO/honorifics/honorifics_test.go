package honorifics_test

import (
	. "intro/honorifics" //to name our package honorifics_test we must import honorifics, this full stop is an alias that makes it so we don't have to do <PACAKAGE>.<FUNCTION> everytime we call a function
	"testing"
)

type testScenario struct { //using a struct to make testing more concise and easier, we can also use many values at the same type
	insertedHonorific string
	expectedResult    bool
}

// Tests standard success of the HonorificType function
func TestHonorificType(t *testing.T) { //test functions mandatorily start with Test, you can then name the rest of the function what you want, the argument must be t *testing.T as well

	//if you want to run tests in parallel, use the t.Parallel() function
	testScenario := []testScenario{
		{"Mister Incredible", true},
		{"Master Splinter", false},
		{"Ronnie", false},
		{"", false},
		{"Miss Incredible", true},
		{"MX DOOM", true},
	}

	for _, scenario := range testScenario {
		actualResult := HonorificType(scenario.insertedHonorific) //actual result of the function
		expectedResult := scenario.expectedResult
		if actualResult != expectedResult {
			t.Errorf("AdressType is not outputting the correct result, expected %t but got %t on name: %s", expectedResult, actualResult, scenario.insertedHonorific)
		} //try to make your test errors as clarifying as possible
	}

} //to run test functions, go to the desired folder (in this case honorifics) and in the terminal run -> go test

//COVERAGE FILES
// to create a more specific coverage file -> go test --coverprofile <GENERATED_FILENAME>.txt on the spefific folder
// to then be able to read this file -> go tool cover --func=<GENERATED_FILENAME>.txt
// to go EVEN FURTHER, to create an html showing you which lines aren't being covered -> go tool cover --html=<GENERATED_FILENAME>.txt
