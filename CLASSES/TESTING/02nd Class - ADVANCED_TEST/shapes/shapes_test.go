package shapes_test

import (
	. "adv/shapes"
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) { //this segregates the rectangle tests from the circle tests
		rect := Rectangle{10, 12}
		expectedResult := float64(120)
		actualResult := rect.Area()

		if expectedResult != actualResult {
			//we can use Faltaf to shut down testing right after the first test fails
			t.Fatalf("Got unexpected area result: %f from Rectangle with dimensions h: %f w: %f, when expected was: %f", actualResult, rect.Height, rect.Width, expectedResult)
		}
	})

	t.Run("Circle", func(t *testing.T) { //segregation again
		circ := Circle{10}
		expectedResult := float64(math.Pi * 100)
		actualResult := circ.Area()

		if expectedResult != actualResult {
			t.Fatalf("Got unexpected area result: %f from Circle with radius: %f, when expected was: %f", actualResult, circ.Radius, expectedResult)
		}
	})
}
