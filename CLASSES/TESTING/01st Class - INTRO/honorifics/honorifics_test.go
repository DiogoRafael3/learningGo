package honorifics

import "testing"

// Tests standard success of the AdressType function
func TestAdressType(t *testing.T) { //test functions mandatorily start with Test, you can then name the rest of the function what you want, the argument must be t *testing.T as well
	testHonorific := "Mister Incredible" //test value

	expectedResult := true //expected result of the function

	actualResult := AdressType(testHonorific) //actual result of the function

	if actualResult != expectedResult {
		t.Errorf("AdressType is not outputting the correct result, expected %t but got %t on name: %s", expectedResult, actualResult, testHonorific)
	} //try to make your test errors as clarifying as possible

} //to run test functions, go to the desired folder (in this case honorifics) and in the terminal run -> go test
